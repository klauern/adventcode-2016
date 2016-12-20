package main

import (
	"reflect"
	"testing"
)

func Test_bot_isOutput(t *testing.T) {
	tests := []struct {
		name string
		b    bot
		want bool
	}{
		{
			"bot 50",
			bot(50),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isOutput(); got != tt.want {
				t.Errorf("bot.isOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bot_isBot(t *testing.T) {
	tests := []struct {
		name string
		b    bot
		want bool
	}{
		{
			"bot 40",
			bot(40),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.isBot(); got != tt.want {
				t.Errorf("bot.isBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_output_isOutput(t *testing.T) {
	tests := []struct {
		name string
		o    output
		want bool
	}{
		{
			"output 50",
			output(50),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.isOutput(); got != tt.want {
				t.Errorf("output.isOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_output_isBot(t *testing.T) {
	tests := []struct {
		name string
		o    output
		want bool
	}{
		{
			"output 40",
			output(40),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.isBot(); got != tt.want {
				t.Errorf("output.isBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compileInstructions(t *testing.T) {
	type args struct {
		instructions []string
	}
	tests := []struct {
		name    string
		args    args
		want    *BotState
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := compileInstructions(tt.args.instructions)
			if (err != nil) != tt.wantErr {
				t.Errorf("compileInstructions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewValue(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name  string
		args  args
		want  bot
		want1 value
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewValue(tt.args.val)
			if got != tt.want {
				t.Errorf("NewValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NewValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewMovement(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name  string
		args  args
		want  bot
		want1 *movement
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewMovement(tt.args.val)
			if got != tt.want {
				t.Errorf("NewMovement() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewMovement() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBotState_IterateCalcs(t *testing.T) {
	type fields struct {
		valueList    map[bot][]value
		movementList map[bot]*movement
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BotState{
				valueList:    tt.fields.valueList,
				movementList: tt.fields.movementList,
			}
			b.IterateCalcs()
		})
	}
}
