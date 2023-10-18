package cmd

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"

	"net/mail"
	"os"
	"path/filepath"
	"time"

	"github.com/gofrs/uuid"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/youmark/pkcs8"
	"https://github.com/AbderrahimBouhdida/AzH/config"
	"https://github.com/AbderrahimBouhdida/AzH/enums"
)

func init() {
	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command{
	Use:          "configure",
	Short:        "Configure AzureHound",
	Run:          configureCmdImpl,
	SilenceUsage: true,
}

func configureCmdImpl(cmd *cobra.Command, args []string) {
	if err := configure(); err != nil {
		exit(fmt.Errorf("failed to configure cobra CLI: %w", err))
	}
}

func configure() error {
	var (
		configFile  = config.ConfigFile.Value().(string)
		configDir   = filepath.Dir(configFile)
		genCert     bool
		genCertPath = filepath.Join(configDir, "cert.pem")
		genKeyPath  = filepath.Join(configDir, "key.pem")
	)

	// Configure Azure connection
	if _, region, err := choose("Region", config.AzRegions, 1); err != nil {
		return err
	} else if tenantId, err := prompt("Directory (tenant) ID", validateGuid, false); err != nil {
		return err
	} else if appId, err := prompt("Application (client) ID", validateGuid, false); err != nil {
		return err
	} else if _, authMethod, err := choose("Authentication Method", enums.AuthMethods(), 0); err != nil {
		return err
	} else {
		config.AzRegion.Set(region)
		config.AzTenant.Set(tenantId)
		config.AzAppId.Set(appId)

		if authMethod == enums.Certificate {
			if genCert = confirm("Generate Certificate and Key", true); genCert {
				if keyPass, err := prompt("Private Key Passphrase (optional)", nil, true); err != nil {
					return err
				} else {
					config.AzCert.Set(genCertPath)
					config.AzKey.Set(genKeyPath)
					config.AzKeyPass.Set(keyPass)
				}
			} else if certPath, err := prompt("Public Certificate Path", validatePem, false); err != nil {
				return err
			} else if keyPath, err := prompt("Private Key Path", validatePem, false); err != nil {
				return err
			} else if keyPass, err := prompt("Private Key Passphrase (optional)", nil, true); err != nil {
				return err
			} else {
				config.AzCert.Set(certPath)
				config.AzKey.Set(keyPath)
				config.AzKeyPass.Set(keyPass)
			}
		} else if authMethod == enums.UsernamePassword {
			if upn, err := prompt("Input Principal Name", validateUserPrincipalName, false); err != nil {
				return err
			} else if password, err := prompt("Input the password", nil, true); err != nil {
				return err
			} else {
				config.AzUsername.Set(upn)
				config.AzPassword.Set(password)
			}
		} else if secret, err := prompt("Client Secret", nil, true); err != nil {
			return err
		} else {
			config.AzSecret.Set(secret)
		}

	}

	// Configure Logging
	if confirm("Setup logging", true) {
		if idx, _, err := choose("Verbosity", verbosityOptions, 1); err != nil {
			return err
		} else if logFile, err := prompt("Log file (optional)", nil, false); err != nil {
			return err
		} else {
			config.VerbosityLevel.Set(idx - 1)
			config.LogFile.Set(logFile)
			config.JsonLogs.Set(confirm("Enable Structured Logs", false))
		}
	}

	// writing the configfile path in the configfile is confusing
	config.ConfigFile.Set(nil)
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return err
	} else if err := viper.WriteConfigAs(configFile); err != nil {
		return err
	} else {
		fmt.Fprintf(os.Stderr, "\nConfiguration written to %s\n", configFile)
	}

	if genCert {
		if cert, key, err := generateCert(config.AzKeyPass.Value().(string)); err != nil {
			return err
		} else if err := os.WriteFile(genCertPath, cert, 0644); err != nil {
			return err
		} else if err := os.WriteFile(genKeyPath, key, 0644); err != nil {
			return err
		} else {
			fmt.Fprintf(os.Stderr, "Key written to %s\n", genKeyPath)
			fmt.Fprintf(os.Stderr, "Certificate written to %s\n", genCertPath)
			fmt.Fprintln(os.Stderr, "\nEnsure certificate is uploaded to your application's client credentials")
		}
	}
	return nil
}

func prompt(label string, validator func(string) error, isSensitive bool) (string, error) {
	p := promptui.Prompt{
		Label:    label,
		Validate: validator,
	}
	if isSensitive {
		p.HideEntered = true
		p.Mask = '*'
	}
	return p.Run()
}

func choose(label string, items []string, pos int) (int, string, error) {
	s := promptui.Select{
		Label:     label,
		CursorPos: pos,
		Items:     items,
		Templates: &promptui.SelectTemplates{
			Selected: fmt.Sprintf(`{{ "%s:" | faint }} {{ . }}`, label),
		},
	}
	return s.Run()
}

func confirm(label string, defaultYes bool) bool {
	p := promptui.Prompt{
		Label:       label,
		HideEntered: true,
		IsConfirm:   true,
	}
	if defaultYes {
		p.Default = "y"
	}
	_, err := p.Run()
	return err == nil
}

func validateGuid(input string) error {
	_, err := uuid.FromString(input)
	return err
}

func validatePem(input string) error {
	if content, err := ioutil.ReadFile(input); err != nil {
		return err
	} else if pemFile, _ := pem.Decode(content); pemFile == nil {
		return fmt.Errorf("Invalid PEM encoded file")
	} else {
		return nil
	}
}

func validateUserPrincipalName(input string) error {
	_, err := mail.ParseAddress(input)
	return err
}

var verbosityOptions = []string{
	"Disabled",
	"Default",
	"Debug",
	"Trace",
}

func generateCert(passphrase string) ([]byte, []byte, error) {
	var (
		cert = &x509.Certificate{
			Subject: pkix.Name{
				CommonName: "azurehound",
			},
			NotBefore: time.Now(),
			NotAfter:  time.Now().AddDate(1, 0, 0),
		}

		certPEM = bytes.Buffer{}
		keyPEM  = bytes.Buffer{}
	)

	// Generate random serial number for certificate
	if serial, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64)); err != nil {
		return nil, nil, err
	} else {
		cert.SerialNumber = serial
	}

	// Generate rsa keys and certificate and encode to PEM files
	if privateKey, err := rsa.GenerateKey(rand.Reader, 4096); err != nil {
		return nil, nil, err
	} else if data, err := x509.CreateCertificate(rand.Reader, cert, cert, &privateKey.PublicKey, privateKey); err != nil {
		return nil, nil, err
	} else if err := pem.Encode(&certPEM, &pem.Block{Type: "CERTIFICATE", Bytes: data}); err != nil {
		return nil, nil, err
	} else if data, err := pkcs8.MarshalPrivateKey(privateKey, []byte(passphrase), nil); err != nil {
		return nil, nil, err
	} else if err := pem.Encode(&keyPEM, &pem.Block{Type: "PRIVATE KEY", Bytes: data}); err != nil {
		return nil, nil, err
	} else {
		return certPEM.Bytes(), keyPEM.Bytes(), nil
	}
}
