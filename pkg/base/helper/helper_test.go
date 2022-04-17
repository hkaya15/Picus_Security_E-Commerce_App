package helper

import (
	"reflect"
	"testing"


	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/cart/model"
	. "github.com/hkaya15/PicusSecurity/Final_Project/pkg/app/category/model"

)

func TestVerifyEMail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "WithoutAtResultFail", args: args{email: "huseyinkayagmail.com"}, want: false},
		{name: "EmptyFieldResultFail", args: args{email: ""}, want: false},
		{name: "WithoutAfterAtSectionResultFail", args: args{email: "huseyinkaya@"}, want: false},
		{name: "FullMailResultSuccess", args: args{email: "huseyinkaya@gmail.com"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyEMail(tt.args.email); got != tt.want {
				t.Errorf("VerifyEMail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		password string
		want     bool
	}{
		{name: "EmptyPasswordResultFail", password: "", want: false},
		{name: "SixCharsResultFail", password: "1abcd.", want: false},
		{name: "SevenCharsWithoutLowerResultFail", password: "1ABCD.E", want: false},
		{name: "SevenCharsWithoutUpperResultFail", password: "1abcd.e", want: false},
		{name: "SevenCharsWithoutNumberResultFail", password: "Xabcd.e", want: false},
		{name: "SevenCharsWithoutPuncResultFail", password: "Xabcd1e", want: false},
		{name: "SevenCharsWithSymbolResultSuccess", password: "Xabc1?e", want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPassword(tt.password); got != tt.want {
				t.Errorf("VerifyEMail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashPassword(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name     string
		password string
		want     error
	}{
		{name: "FullPasswordResultSuccess", password: "123", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := HashPassword(tt.password); got != tt.want {
				t.Errorf("VerifyEMail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "WithHashResultSuccess", args: args{password: "123", hash: "arc"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		clist CategoryList
		c     Category
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "SameListResultTrue", args: args{clist: []Category{
			{CategoryID: "1", CategoryName: "ABC", IconURL: "abc"},
		},
			c: Category{CategoryID: "1", CategoryName: "ABC", IconURL: "abc"},
		}, want: true},
		{name: "DifferentListResultFalse", args: args{clist: []Category{
			{CategoryID: "2", CategoryName: "ABCD", IconURL: "abcd"},
		},
			c: Category{CategoryID: "1", CategoryName: "ABC", IconURL: "abc"},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.clist, tt.args.c); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateValues(t *testing.T) {
	type args struct {
		cart      Cart
		cartitems []CartsItem
	}
	tests := []struct {
		name string
		args args
		want *Cart
	}{
		{name: "UpdateWithSameValuesResultSuccess",args: args{
			cart: Cart{UserID: "1"},
			cartitems: []CartsItem{
				{TotalPrice: 25},
				{TotalPrice: 25},
			},
		},
	want: &Cart{CartTotalPrice: 50, CartLength: 2},},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateValues(tt.args.cart, tt.args.cartitems); reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
