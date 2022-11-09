# loofah

Print a time-based 2FA code from a supplied 2FA secret key.

## Install

```
go install github.com/nishanths/loofah@latest
```

## Differences from github.com/rsc/2fa

You might prefer this program to [github.com/rsc/2fa][1] if you want to
manually provide the 2FA secret key yourself each time when you need a code
(e.g., you may have the secret key stored in your own password manager).

This program only supports obtaining time-based (TOTP) authentication codes.

## Usage

```
usage: loofah [-7] [-8] [-c]
```

Example:

```
% loofah
enter 2fa key: nzxxiidbebvwk6jb
852415
%
```

In this example the text `nzxxiidbebvwk6jb` is a 2FA secret key.

On many apps, during 2FA setup, the 2FA secret key can be viewed by
selecting an option like "Can't scan QR code?" or "Enter this text code
instead." Save this secret key in your password manager. Provide the
secret key to `loofah` when you need a code.

By default `loofah` prints a 6-digit code. Use the `-7` flag or the `-8`
flag to print a 7-digit or 8-digit code. Specify the `-c` flag to also
copy the code to the clipboard.

The time-based authentication codes are derived from a hash of the key
and the current time, so it is important that the system clock have at
least one-minute accuracy.

## License

See the `LICENSE` file.

Code in `loofah.go` is adapted from [github.com/rsc/2fa][1]; see license
information in `loofah.go`.

[1]: https://github.com/rsc/2fa
