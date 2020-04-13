package jwt

// https://github.com/nats-io/jwt

/*
Audience	表示JWT的受众
ExpiresAt	失效时间
Id	签发编号
IssuedAt	签发时间
Issuer	签发人
NotBefore	生效时间
Subject	主题

Subject:   "login",           // 主题
"data": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxMjMiLCJleHAiOjE1NjQ3OTY0MzksImp0aSI6Ilx1MDAwMCIsImlhdCI6MTU2NDc5NjQxOSwiaXNzIjoiZ2luIGhlbGxvIiwibmJmIjoxNTY0Nzk2NDE5LCJzdWIiOiJsb2dpbiJ9.CpacmfBSMgmK2TgrT-KwNB60bsvwgyryGQ0pWZr8laU"
*/