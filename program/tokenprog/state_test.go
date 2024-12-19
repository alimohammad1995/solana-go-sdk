package tokenprog

import (
	"testing"

	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/pkg/pointer"
	"github.com/stretchr/testify/assert"
)

func TestAccountFromData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    TokenAccount
		wantErr error
	}{
		{
			args: args{
				data: []byte{105, 145, 9, 101, 129, 184, 46, 130, 176, 132, 102, 98, 17, 241, 215, 189, 90, 219, 106, 196, 196, 121, 174, 243, 65, 40, 132, 7, 252, 112, 238, 112, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 0, 186, 69, 61, 244, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: TokenAccount{
				Mint:            common.PublicKeyFromString("8765cK2Vucsic6NA5nm4cfkrCzusaFVqBf6Pk31tGkXH"),
				Owner:           common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				Amount:          1049000000000,
				Delegate:        nil,
				State:           TokenAccountStateInitialized,
				IsNative:        nil,
				DelegatedAmount: 0,
				CloseAuthority:  nil,
			},
			wantErr: nil,
		},
		{
			args: args{
				data: []byte{0x6, 0x9b, 0x88, 0x57, 0xfe, 0xab, 0x81, 0x84, 0xfb, 0x68, 0x7f, 0x63, 0x46, 0x18, 0xc0, 0x35, 0xda, 0xc4, 0x39, 0xdc, 0x1a, 0xeb, 0x3b, 0x55, 0x98, 0xa0, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x10, 0x96, 0x59, 0x17, 0x5e, 0x7c, 0x64, 0x33, 0x21, 0xa5, 0xed, 0x46, 0x42, 0xa0, 0x27, 0xb0, 0xab, 0xd9, 0x7b, 0x8d, 0xd9, 0x7a, 0xd1, 0xbc, 0xc6, 0xdc, 0x64, 0x71, 0x38, 0x6c, 0xcd, 0xdc, 0x10, 0x76, 0x16, 0x77, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, 0xf0, 0x1d, 0x1f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: TokenAccount{
				Mint:            common.NativeMint,
				Owner:           common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
				Amount:          1997960720,
				Delegate:        nil,
				State:           TokenAccountStateInitialized,
				IsNative:        pointer.Get[uint64](2039280),
				DelegatedAmount: 0,
				CloseAuthority:  nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TokenAccountFromData(tt.args.data)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMintAccountFromData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    MintAccount
		wantErr error
	}{
		{
			args: args{
				data: []byte{0x1, 0x0, 0x0, 0x0, 0x10, 0x96, 0x59, 0x17, 0x5e, 0x7c, 0x64, 0x33, 0x21, 0xa5, 0xed, 0x46, 0x42, 0xa0, 0x27, 0xb0, 0xab, 0xd9, 0x7b, 0x8d, 0xd9, 0x7a, 0xd1, 0xbc, 0xc6, 0xdc, 0x64, 0x71, 0x38, 0x6c, 0xcd, 0xdc, 0x0, 0xe4, 0xb, 0x54, 0x2, 0x0, 0x0, 0x0, 0x9, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: MintAccount{
				MintAuthority:   pointer.Get[common.PublicKey](common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ")),
				Supply:          10000000000,
				Decimals:        9,
				IsInitialized:   true,
				FreezeAuthority: nil,
			},
			wantErr: nil,
		},
		{
			args: args{
				data: []byte{0x1, 0x0, 0x0, 0x0, 0x10, 0x96, 0x59, 0x17, 0x5e, 0x7c, 0x64, 0x33, 0x21, 0xa5, 0xed, 0x46, 0x42, 0xa0, 0x27, 0xb0, 0xab, 0xd9, 0x7b, 0x8d, 0xd9, 0x7a, 0xd1, 0xbc, 0xc6, 0xdc, 0x64, 0x71, 0x38, 0x6c, 0xcd, 0xdc, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa, 0x1, 0x1, 0x0, 0x0, 0x0, 0x10, 0x96, 0x59, 0x17, 0x5e, 0x7c, 0x64, 0x33, 0x21, 0xa5, 0xed, 0x46, 0x42, 0xa0, 0x27, 0xb0, 0xab, 0xd9, 0x7b, 0x8d, 0xd9, 0x7a, 0xd1, 0xbc, 0xc6, 0xdc, 0x64, 0x71, 0x38, 0x6c, 0xcd, 0xdc},
			},
			want: MintAccount{
				MintAuthority:   pointer.Get[common.PublicKey](common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ")),
				Supply:          0,
				Decimals:        10,
				IsInitialized:   true,
				FreezeAuthority: pointer.Get[common.PublicKey](common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ")),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MintAccountFromData(tt.args.data)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMultisigAccountFromData(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    MultisigAccount
		wantErr error
	}{
		{
			name: "1 / 3",
			args: args{
				data: []byte{0x1, 0x3, 0x1, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x70, 0xf9, 0xab, 0xec, 0xd2, 0x24, 0x46, 0x4b, 0x71, 0x86, 0x3d, 0xc9, 0x1f, 0x19, 0xc4, 0x6f, 0x87, 0x48, 0x52, 0x97, 0x62, 0x3e, 0x60, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x78, 0x5f, 0xcb, 0xec, 0x99, 0xc2, 0xb, 0x73, 0x3d, 0x6b, 0xc1, 0xf2, 0xbf, 0xd6, 0x18, 0x39, 0xc8, 0xbe, 0x52, 0xc8, 0xc4, 0x93, 0x70, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x7f, 0xc5, 0xeb, 0xec, 0x61, 0x5f, 0xd0, 0x9b, 0x9, 0x51, 0x46, 0x1c, 0x60, 0x92, 0x6c, 0x4, 0xa, 0x34, 0x52, 0xfa, 0x26, 0xe8, 0x80, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: MultisigAccount{
				M:             1,
				N:             3,
				IsInitialized: true,
				Signers: []common.PublicKey{
					common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner3111111111111111111111111111111111111"),
				},
			},
			wantErr: nil,
		},
		{
			name: "3 / 5",
			args: args{
				data: []byte{0x3, 0x5, 0x1, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x70, 0xf9, 0xab, 0xec, 0xd2, 0x24, 0x46, 0x4b, 0x71, 0x86, 0x3d, 0xc9, 0x1f, 0x19, 0xc4, 0x6f, 0x87, 0x48, 0x52, 0x97, 0x62, 0x3e, 0x60, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x78, 0x5f, 0xcb, 0xec, 0x99, 0xc2, 0xb, 0x73, 0x3d, 0x6b, 0xc1, 0xf2, 0xbf, 0xd6, 0x18, 0x39, 0xc8, 0xbe, 0x52, 0xc8, 0xc4, 0x93, 0x70, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x7f, 0xc5, 0xeb, 0xec, 0x61, 0x5f, 0xd0, 0x9b, 0x9, 0x51, 0x46, 0x1c, 0x60, 0x92, 0x6c, 0x4, 0xa, 0x34, 0x52, 0xfa, 0x26, 0xe8, 0x80, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x87, 0x2c, 0xb, 0xec, 0x28, 0xfd, 0x95, 0xc2, 0xd5, 0x36, 0xca, 0x46, 0x1, 0x4e, 0xbf, 0xce, 0x4b, 0xaa, 0x53, 0x2b, 0x89, 0x3d, 0x90, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x8e, 0x92, 0x2b, 0xeb, 0xf0, 0x9b, 0x5a, 0xea, 0xa1, 0x1c, 0x4e, 0x6f, 0xa2, 0xb, 0x13, 0x98, 0x8d, 0x20, 0x53, 0x5c, 0xeb, 0x92, 0xa0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: MultisigAccount{
				M:             3,
				N:             5,
				IsInitialized: true,
				Signers: []common.PublicKey{
					common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner3111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner4111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner5111111111111111111111111111111111111"),
				},
			},
			wantErr: nil,
		},
		{
			name: "10 / 11",
			args: args{
				data: []byte{0xa, 0xb, 0x1, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x70, 0xf9, 0xab, 0xec, 0xd2, 0x24, 0x46, 0x4b, 0x71, 0x86, 0x3d, 0xc9, 0x1f, 0x19, 0xc4, 0x6f, 0x87, 0x48, 0x52, 0x97, 0x62, 0x3e, 0x60, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x78, 0x5f, 0xcb, 0xec, 0x99, 0xc2, 0xb, 0x73, 0x3d, 0x6b, 0xc1, 0xf2, 0xbf, 0xd6, 0x18, 0x39, 0xc8, 0xbe, 0x52, 0xc8, 0xc4, 0x93, 0x70, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x7f, 0xc5, 0xeb, 0xec, 0x61, 0x5f, 0xd0, 0x9b, 0x9, 0x51, 0x46, 0x1c, 0x60, 0x92, 0x6c, 0x4, 0xa, 0x34, 0x52, 0xfa, 0x26, 0xe8, 0x80, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x87, 0x2c, 0xb, 0xec, 0x28, 0xfd, 0x95, 0xc2, 0xd5, 0x36, 0xca, 0x46, 0x1, 0x4e, 0xbf, 0xce, 0x4b, 0xaa, 0x53, 0x2b, 0x89, 0x3d, 0x90, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x8e, 0x92, 0x2b, 0xeb, 0xf0, 0x9b, 0x5a, 0xea, 0xa1, 0x1c, 0x4e, 0x6f, 0xa2, 0xb, 0x13, 0x98, 0x8d, 0x20, 0x53, 0x5c, 0xeb, 0x92, 0xa0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x95, 0xf8, 0x4b, 0xeb, 0xb8, 0x39, 0x20, 0x12, 0x6d, 0x1, 0xd2, 0x99, 0x42, 0xc7, 0x67, 0x62, 0xce, 0x96, 0x53, 0x8e, 0x4d, 0xe7, 0xb0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0x9d, 0x5e, 0x6b, 0xeb, 0x7f, 0xd6, 0xe5, 0x3a, 0x38, 0xe7, 0x56, 0xc2, 0xe3, 0x83, 0xbb, 0x2d, 0x10, 0xc, 0x53, 0xbf, 0xb0, 0x3c, 0xc0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0xa4, 0xc4, 0x8b, 0xeb, 0x47, 0x74, 0xaa, 0x62, 0x4, 0xcc, 0xda, 0xec, 0x84, 0x40, 0xe, 0xf7, 0x51, 0x82, 0x53, 0xf1, 0x12, 0x91, 0xd0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0xac, 0x2a, 0xab, 0xeb, 0xf, 0x12, 0x6f, 0x89, 0xd0, 0xb2, 0x5f, 0x16, 0x24, 0xfc, 0x62, 0xc1, 0x92, 0xf8, 0x54, 0x22, 0x74, 0xe6, 0xe0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0xb3, 0x90, 0xcb, 0xea, 0xd6, 0xb0, 0x34, 0xb1, 0x9c, 0x97, 0xe3, 0x3f, 0xc5, 0xb8, 0xb6, 0x8b, 0xd4, 0x6e, 0x54, 0x53, 0xd7, 0x3b, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x4b, 0xbe, 0x48, 0xba, 0xf6, 0xeb, 0xea, 0x9e, 0x4d, 0xf9, 0xd9, 0x68, 0x7d, 0x67, 0x69, 0x66, 0x75, 0xa, 0x56, 0x15, 0xe4, 0x54, 0x85, 0x39, 0x91, 0x0, 0x0, 0x0, 0x0, 0x0},
			},
			want: MultisigAccount{
				M:             10,
				N:             11,
				IsInitialized: true,
				Signers: []common.PublicKey{
					common.PublicKeyFromString("S1gner1111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner2111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner3111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner4111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner5111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner6111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner7111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner8111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gner9111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gnerA111111111111111111111111111111111111"),
					common.PublicKeyFromString("S1gnerB111111111111111111111111111111111111"),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MultisigAccountFromData(tt.args.data)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeserializeTokenAccount(t *testing.T) {
	type args struct {
		data         []byte
		accountOwner common.PublicKey
	}
	tests := []struct {
		name    string
		args    args
		want    TokenAccount
		wantErr error
	}{
		{
			args: args{
				data:         []byte{105, 145, 9, 101, 129, 184, 46, 130, 176, 132, 102, 98, 17, 241, 215, 189, 90, 219, 106, 196, 196, 121, 174, 243, 65, 40, 132, 7, 252, 112, 238, 112, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 0, 186, 69, 61, 244, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				accountOwner: common.SystemProgramID,
			},
			want:    TokenAccount{},
			wantErr: ErrInvalidAccountOwner,
		},
		{
			args: args{
				data:         []byte{1, 105, 145, 9, 101, 129, 184, 46, 130, 176, 132, 102, 98, 17, 241, 215, 189, 90, 219, 106, 196, 196, 121, 174, 243, 65, 40, 132, 7, 252, 112, 238, 112, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 0, 186, 69, 61, 244, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				accountOwner: common.TokenProgramID,
			},
			want:    TokenAccount{},
			wantErr: ErrInvalidAccountDataSize,
		},
		{
			args: args{
				data:         []byte{145, 9, 101, 129, 184, 46, 130, 176, 132, 102, 98, 17, 241, 215, 189, 90, 219, 106, 196, 196, 121, 174, 243, 65, 40, 132, 7, 252, 112, 238, 112, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 0, 186, 69, 61, 244, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				accountOwner: common.TokenProgramID,
			},
			want:    TokenAccount{},
			wantErr: ErrInvalidAccountDataSize,
		},
		{
			args: args{
				data:         []byte{105, 145, 9, 101, 129, 184, 46, 130, 176, 132, 102, 98, 17, 241, 215, 189, 90, 219, 106, 196, 196, 121, 174, 243, 65, 40, 132, 7, 252, 112, 238, 112, 206, 211, 135, 230, 195, 111, 87, 254, 147, 239, 143, 81, 110, 159, 49, 140, 109, 137, 224, 197, 24, 49, 223, 61, 123, 8, 78, 109, 110, 136, 228, 240, 0, 186, 69, 61, 244, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				accountOwner: common.TokenProgramID,
			},
			want: TokenAccount{
				Mint:            common.PublicKeyFromString("8765cK2Vucsic6NA5nm4cfkrCzusaFVqBf6Pk31tGkXH"),
				Owner:           common.PublicKeyFromString("EvN4kgKmCmYzdbd5kL8Q8YgkUW5RoqMTpBczrfLExtx7"),
				Amount:          1049000000000,
				Delegate:        nil,
				State:           TokenAccountStateInitialized,
				IsNative:        nil,
				DelegatedAmount: 0,
				CloseAuthority:  nil,
			},
			wantErr: nil,
		},
		{
			args: args{
				data:         []byte{0x6, 0x9b, 0x88, 0x57, 0xfe, 0xab, 0x81, 0x84, 0xfb, 0x68, 0x7f, 0x63, 0x46, 0x18, 0xc0, 0x35, 0xda, 0xc4, 0x39, 0xdc, 0x1a, 0xeb, 0x3b, 0x55, 0x98, 0xa0, 0xf0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x10, 0x96, 0x59, 0x17, 0x5e, 0x7c, 0x64, 0x33, 0x21, 0xa5, 0xed, 0x46, 0x42, 0xa0, 0x27, 0xb0, 0xab, 0xd9, 0x7b, 0x8d, 0xd9, 0x7a, 0xd1, 0xbc, 0xc6, 0xdc, 0x64, 0x71, 0x38, 0x6c, 0xcd, 0xdc, 0x10, 0x76, 0x16, 0x77, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x1, 0x0, 0x0, 0x0, 0xf0, 0x1d, 0x1f, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				accountOwner: common.TokenProgramID,
			},
			want: TokenAccount{
				Mint:            common.NativeMint,
				Owner:           common.PublicKeyFromString("27kVX7JpPZ1bsrSckbR76mV6GeRqtrjoddubfg2zBpHZ"),
				Amount:          1997960720,
				Delegate:        nil,
				State:           TokenAccountStateInitialized,
				IsNative:        pointer.Get[uint64](2039280),
				DelegatedAmount: 0,
				CloseAuthority:  nil,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeserializeTokenAccount(tt.args.data, tt.args.accountOwner)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
