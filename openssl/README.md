# RSA Signature Verification with OpenSSL

This document explains how to verify an RSA signature using OpenSSL command-line tools.

## Files Overview

- `message.txt` - The original message that was signed
- `public_key.pem` - The RSA public key used for verification
- `signature.bin` - The binary signature file (to be created from base64 text)

## Step 1: Convert Base64 Signature to Binary

If you have a base64-encoded signature string like:
```
Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=
```

Convert it to binary file:
```bash
echo "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=" | base64 -d > signature.bin
```

Or save to a file first:
```bash
# Save base64 to file
echo "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=" > signature.b64

# Convert to binary
base64 -d signature.b64 > signature.bin
```

## Step 2: Verify the Signature

Once you have the binary signature file, verify it:

```bash
openssl dgst -sha256 -verify public_key.pem -signature signature.bin message.txt
```

This command:
- Uses SHA-256 as the digest algorithm (adjust if different algorithm was used)
- Reads the public key from `public_key.pem`
- Reads the binary signature from `signature.bin`
- Verifies the signature against `message.txt`

Expected output if verification succeeds:
```
Verified OK
```

## Alternative Method: One-line Verification from Base64

You can verify directly from base64 without creating a binary file:
```bash
echo "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=" | base64 -d | openssl dgst -sha256 -verify public_key.pem -signature /dev/stdin message.txt
```

## Common Hash Algorithms

If SHA-256 doesn't work, try these common alternatives:
- `-sha1` - SHA-1 (less secure, but still used in legacy systems)
- `-sha512` - SHA-512
- `-md5` - MD5 (deprecated, avoid for new implementations)

## Troubleshooting

1. **"Verification Failure"** - This means the signature doesn't match. Possible causes:
   - Wrong hash algorithm
   - Message was modified after signing
   - Wrong public key
   - Corrupted signature

2. **Key format issues** - Ensure the public key is in PEM format (starts with `-----BEGIN PUBLIC KEY-----`)

3. **Binary vs Text** - Ensure `message.txt` has the exact same content and encoding as when it was signed

## Example Verification Script

```bash
#!/bin/bash

# Verify binary signature
echo "Verifying signature..."
if openssl dgst -sha256 -verify public_key.pem -signature signature.bin message.txt; then
    echo "Signature is valid!"
else
    echo "Signature verification failed!"
    exit 1
fi
```

## Security Notes

- Always verify signatures using the correct public key
- Ensure the message hasn't been tampered with
- Use secure hash algorithms (SHA-256 or higher)
- Keep public keys in a trusted location