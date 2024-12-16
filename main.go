package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"

	"github.com/spf13/cobra"
)

type PasswordOptions struct {
	Length  int
	Lower   bool
	Upper   bool
	Digits  bool
	Special bool
}

var pwOpts PasswordOptions
var rootCmd = &cobra.Command{
	Use:   "pwgen",
	Short: "Generate a random password",
	Long: `Generate a random password.
By default, the password will be 21 characters long and include lowercase letters, uppercase letters, and digits.`,

	Run: func(cmd *cobra.Command, args []string) {
		password, err := generatePassword(pwOpts)
		if err != nil {
			fmt.Println("Error generating password:", err)
			os.Exit(1)
		}

		fmt.Printf("%s", password)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&pwOpts.Length, "length", "L", 21, "Length of the password")
	rootCmd.Flags().BoolVarP(&pwOpts.Lower, "lower", "l", false, "Include lowercase characters")
	rootCmd.Flags().BoolVarP(&pwOpts.Upper, "upper", "u", false, "Include uppercase characters")
	rootCmd.Flags().BoolVarP(&pwOpts.Digits, "digits", "d", false, "Include digits")
	rootCmd.Flags().BoolVarP(&pwOpts.Special, "special", "s", false, "Include special characters")
}

func generatePassword(opts PasswordOptions) (string, error) {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	special := "!@#$%^&*()_-+=<>?"

	if opts.Length < 1 {
		return "", fmt.Errorf("length must be greater than 0")
	}

	chars := ""
	if opts.Lower {
		chars += lower
	}
	if opts.Upper {
		chars += upper
	}
	if opts.Digits {
		chars += digits
	}
	if opts.Special {
		chars += special
	}

	if chars == "" {
		chars = lower + upper + digits
	}

	password := ""
	for i := 0; i < opts.Length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		password += string(chars[index.Int64()])
	}

	return password, nil
}

func main() {
	rootCmd.Execute()
}
