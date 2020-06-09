# totp-proxy
protect your website behind time-based one-time password encryption.

```flow
          +-----------+      +-----------+
          |   T-OTP   |      |           |
          |           |      |           |
          |           |      |           |
          |           |      |           |
    token |           |      |           |
   +------>   proxy   +------>  website  |
          |           |      |           |
          |           |      |           |
          |           |      |           |
          |           |      |           |
          |           |      |           |
          +-----------+      +-----------+
```

## Usage

totp is easiest using and safest encryption.

1. run totp-proxy behind your sensitive website
2. visit proxy server with time-based token

### command usage

```
$ totp-proxy help                      
NAME:
   totp-proxy - using time-based one-time password protect your website

USAGE:
   totp-proxy [global options] command [command options] upstream

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --secret value, -s value    base32 encoding totp secret (default: "NB2HI4B2F4XXM2LOMNSW45DDOVUS4Y3O") [$SECRET]
   --interval value, -i value  totp token refresh interval (default: 1800) [$INTERVAL]
   --help, -h                  show help (default: false)
```

### server

run command below and you will get output.

```bash
$ totp-proxy {your_web_site} 
2020/06/09 21:44:45 using secret: NB2HI4B2F4XXM2LOMNSW45DDOVUS4Y3O
2020/06/09 21:44:45 using interval: 1800
```

### guest access

visit `http://proxy/` on guest's browser, and get `401` response.

### security access

access your website with token

#### token

1. open totp online generator like https://totp.danhersam.com
2. copy and paste secret into secret input (save your secret properly)
3. type interval into interval input
4. your get 6 digits time-based 

> using T-OTP with Google Authenticator or 1Password

visit `http://proxy/?x-totp-token={token}` on your browser,
then you get your treasure

## T-OTP

token is temporary key 

**!!! SAVE YOUR SECRET PROPERLY**
