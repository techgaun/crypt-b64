# crypt-b64
> Quick simple tool to encode crypt based base64 encoded text to standard base64 encoding

Crypt's encoding uses `./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz` as the charset
and base64 encoding (as we know & use a lot) uses `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/` as the charset.

## Example

Given hash: `$6$HuxJVXbQPLPnqRB$DHCTcDd7xa2pB1l/e64Zg8kRmRuV0EChEoftE5rf3.L2B168CA2ir5/EGIBkwLw1Un3E8.43DXe4wwtjqLrxi1`

```shell
./crypt-b64 'DHCTcDd7xa2pB1l/e64Zg8kRmRuV0EChEoftE5rf3.L2B168CA2ir5/EGIBkwLw1Un3E8.43DXe4wwtjqLrxi1'
```

## Why

Ikr. Was needed for custom tooling.
