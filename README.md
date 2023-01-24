# Gimme Vault
Configurable Golang application to login to AWS CLI

## Features
- Easy to configure
- Supports multiple profiles
- Reuses token to avoid multiple logins

## Requirements
- Connect to the appropriate VPN
- [Golang](https://go.dev/)
- That's it!!

## Build

```bash
go build .
```

## How to use

#### Configure
`./gimme-vault configure [--profile <profile_name>]`

#### Login
`./gimme-vault login [--profile <profile_name>]`

> **profile** flag defaults to: *default*

### Sample configuration

https://some_url:8200/v1/\<resource\>/aws_account/\<action\>
```
username = your_ldap_username
awsAccount = aws_account        // 123456789012
region = aws_region             // us-west-2
url = vault_addr                // eg: https://some_url:8200
version = v1                    // Check wiki CURL
resource = ...                  // Check wiki CURL
action = ...                    // Check wiki CURL
```


## Side notes

* Requests AWS Credentials for "240m" (240 minutes)
* Timeout errors could be due to VPN connectivity
* Configuration file is saved in home directory by default as `.gimme-vault.yaml`