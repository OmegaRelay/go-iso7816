// SPDX-FileCopyrightText: 2023 Steffen Vogel <post@steffenvogel.de>
// SPDX-License-Identifier: Apache-2.0

package yubikey

import (
	"errors"

	iso "cunicu.li/go-iso7816"
)

var ErrInvalidResponseLength = errors.New("invalid response length")

const (
	// https://docs.yubico.com/yesdk/users-manual/application-otp/otp-commands.html
	InsOTP        iso.Instruction = 0x01 // Most commands of the OTP applet use this value
	InsReadStatus iso.Instruction = 0x03
)

// func withYubicoApplet(cb func(tx *Transaction) (bool, error)) Filter {
// 	return func(name string, card *Card) (bool, error) {
// 		// Matching against the name first saves us from connecting to the card
// 		if match, err := IsYubikey(name, card); err != nil {
// 			return false, err
// 		} else if !match {
// 			return false, nil
// 		}

// 		if card == nil {
// 			return false, ErrOpen
// 		}

// 		tx, err := card.NewTransaction()
// 		if err != nil {
// 			return false, fmt.Errorf("failed to begin transaction: %w", err)
// 		}
// 		defer tx.Close()

// 		if _, err := tx.Select(AidYubiKeyOTP); err != nil {
// 			return false, nil
// 		}

// 		return cb(tx)
// 	}
// }
