# loofah

Generate a 6, 7, or 8-digit 2FA code from a 2FA secret key.

## Install

```
go install github.com/nishanths/loofah@latest
```

## Differences from github.com/rsc/2fa

You may prefer this program to [github.com/rsc/2fa][1] if you want to manually
provide the 2FA secret key yourself each time when you need a 2FA code (e.g.,
if you have the 2FA secret key stored in a separate password manager).

This program only supports obtaining time-based (TOTP) authentication codes.

## Usage

```
usage: loofah [-7] [-8]
```

Example:

```
% loofah
enter 2fa key: nzxxiidbebvwk6jb
852415
%
```

In this example, the input text `nzxxiidbebvwk6jb` is a 2FA secret key.

On many apps, during 2FA setup, the 2FA secret key can be viewed by choosing
"Can't scan QR code?" or "Enter this text code instead". Save these 2FA secret
keys in your password manager. Then provide them to `loofah` when you need a
2FA login code.

By default a 6-digit 2FA code is printed. Use the `-7` flag or `-8` flag to
produce a 7 or 8-digit code.

The time-based authentication codes are derived from a hash of the key and the
current time, so it is important that the system clock have at least
one-minute accuracy.

## License

See the `LICENSE` file.

Code in `loofah.go` is adapted from [github.com/rsc/2fa][1]; see license
information in `loofah.go`.

[1]: https://github.com/rsc/2fa
