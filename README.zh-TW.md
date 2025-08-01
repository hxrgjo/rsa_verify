# 使用 OpenSSL 進行 RSA 簽章驗證

本文件說明如何使用 OpenSSL 命令列工具來驗證 RSA 簽章。

## 檔案概覽

- `message.txt` - 被簽章的原始訊息
- `public_key.pem` - 用於驗證的 RSA 公鑰
- `signature.bin` - 二進位簽章檔案（將從 base64 文字建立）

## 步驟一：將 Base64 簽章轉換為二進位檔案

如果您有一段 base64 編碼的簽章字串，例如：
```
Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=
```

將其轉換為二進位檔案：
```bash
echo "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=" | base64 -d > signature.bin
```

或者先存成檔案：
```bash
# 儲存 base64 到檔案
echo "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=" > signature.b64

# 轉換為二進位
base64 -d signature.b64 > signature.bin
```

## 步驟二：驗證簽章

當您擁有二進位簽章檔案後，進行驗證：

```bash
openssl dgst -sha256 -verify public_key.pem -signature signature.bin message.txt
```

這個指令：
- 使用 SHA-256 作為摘要演算法（如果使用不同演算法請調整）
- 從 `public_key.pem` 讀取公鑰
- 從 `signature.bin` 讀取二進位簽章
- 驗證簽章是否與 `message.txt` 相符

驗證成功時的預期輸出：
```
Verified OK
```

## 替代方法：從 Base64 直接驗證（單行指令）

您也可以直接從 base64 驗證而不建立二進位檔案：
```bash
echo "Kg5aEcTdxcI772NZItxg6ZN27nGp4xLxxQqOsbONILxnrA/vFtqZKRnxrIp+/QkvYQR0Fc7uNiFdZLJyt4+qesVCov2y+vKVfclpxaKZ65nwrdKCP8yWJ8fJuko+t4UUtON8f/6yNshk0J3LF/9vZeZuu8hOcigNuPysWhOqDLE=" | base64 -d | openssl dgst -sha256 -verify public_key.pem -signature /dev/stdin message.txt
```

## 常見雜湊演算法

如果 SHA-256 無法運作，請嘗試以下常見替代方案：
- `-sha1` - SHA-1（安全性較低，但仍用於舊系統）
- `-sha512` - SHA-512
- `-md5` - MD5（已棄用，避免用於新實作）

## 疑難排解

1. **「驗證失敗」** - 這表示簽章不符。可能原因：
   - 錯誤的雜湊演算法
   - 簽章後訊息被修改
   - 錯誤的公鑰
   - 簽章損毀

2. **金鑰格式問題** - 確保公鑰為 PEM 格式（以 `-----BEGIN PUBLIC KEY-----` 開頭）

3. **二進位與文字** - 確保 `message.txt` 的內容和編碼與簽章時完全相同

## 驗證腳本範例

```bash
#!/bin/bash

# 驗證二進位簽章
echo "正在驗證簽章..."
if openssl dgst -sha256 -verify public_key.pem -signature signature.bin message.txt; then
    echo "簽章有效！"
else
    echo "簽章驗證失敗！"
    exit 1
fi
```

## 安全注意事項

- 務必使用正確的公鑰驗證簽章
- 確保訊息未被竄改
- 使用安全的雜湊演算法（SHA-256 或更高）
- 將公鑰保存在受信任的位置