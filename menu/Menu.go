package cmenu

import "github.com/manifoldco/promptui"

type challenge struct {
	Name        string
	Description string
}

func GetPrompt() promptui.Select {
	return promptui.Select{
		Label:     "Select challenge:",
		Items:     GetChallenges(),
		Templates: GetTemplates(),
		Size:      4,
	}
}

// get challenges objects
func GetChallenges() []challenge {
	return []challenge{
		{Name: "C1 - Convert hex to base64", Description: `The string: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
Should produce: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t`},
		{Name: "C2 - Fixed XOR", Description: `Write a function that takes two equal-length buffers and produces their XOR combination.
If your function works properly, then when you feed it the string: 1c0111001f010100061a024b53535009181c
... after hex decoding, and when XOR'd against: 686974207468652062756c6c277320657965
... should produce: 746865206b696420646f6e277420706c6179`},
		{Name: "C3 - Single-byte XOR cipher", Description: `The hex encoded string: 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message. `},
		{Name: "C4 - Detect single-character XOR", Description: `The string: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
Should produce: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t`},
	}
}

// get templates objects
func GetTemplates() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . | red }}",
		Active:   "👉 {{ .Name | cyan }}",
		Inactive: "   {{ .Name | blue }}",
		Selected: "👉 {{ .Name | green }}",
		Details: `
{{ "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁" | green }} 
{{ "▏" | green }}          Details          {{ "▕" | green }} 
{{ "▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔" | green }}
{{ "Name: " | cyan | faint }}	{{ .Name }}
{{ "Description: " | red | faint }}
{{ .Description }}`,
	}
}
