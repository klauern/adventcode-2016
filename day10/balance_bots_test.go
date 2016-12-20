package main

import (
	"reflect"
	"testing"

	"github.com/klauern/adventcode-2016/helpers"
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
		{
			"value 5 goes to bot 2",
			args{"value 5 goes to bot 2"},
			bot(2),
			value(5),
		},
		{
			"value 3 goes to bot 1",
			args{"value 3 goes to bot 1"},
			bot(1),
			value(3),
		},
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
		{
			"bot 2 gives low to bot 1 and high to bot 0",
			args{"bot 2 gives low to bot 1 and high to bot 0"},
			bot{id: botID(2), vals: []value{value(1)}},
			&movement{lowTo: bot(1), highTo: bot(0)},
		},
		{
			"bot 1 gives low to output 1 and high to bot 0",
			args{"bot 1 gives low to output 1 and high to bot 0"},
			bot(1),
			&movement{lowTo: output(1), highTo: bot(0)},
		},
		{
			"bot 0 gives low to output 2 and high to output 0",
			args{"bot 0 gives low to output 2 and high to output 0"},
			bot(0),
			&movement{lowTo: output(2), highTo: output(0)},
		},
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

func TestValueCount(t *testing.T) {
	file := helpers.MustScanFile("input")
	for file.Scan() {
		line := file.Text()
	}
}

func Test_output_canReceive(t *testing.T) {
	type fields struct {
		id   outputID
		vals []value
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"output (true)",
			fields{outputID(50), []vals{}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := output{
				id:   tt.fields.id,
				vals: tt.fields.vals,
			}
			if got := o.canReceive(); got != tt.want {
				t.Errorf("output.canReceive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bot_canReceive(t *testing.T) {
	type fields struct {
		id   botID
		vals []value
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bot{
				id:   tt.fields.id,
				vals: tt.fields.vals,
			}
			if got := b.canReceive(); got != tt.want {
				t.Errorf("bot.canReceive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_output_receiveValue(t *testing.T) {
	type fields struct {
		id   outputID
		vals []value
	}
	type args struct {
		val value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := output{
				id:   tt.fields.id,
				vals: tt.fields.vals,
			}
			o.receiveValue(tt.args.val)
		})
	}
}

func Test_bot_receiveValue(t *testing.T) {
	type fields struct {
		id   botID
		vals []value
	}
	type args struct {
		val value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bot{
				id:   tt.fields.id,
				vals: tt.fields.vals,
			}
			b.receiveValue(tt.args.val)
		})
	}
}

func TestBotState_move(t *testing.T) {
	type fields struct {
		bots         map[botID]*bot
		movementList map[botID]*movement
		outputList   map[int][]value
	}
	type args struct {
		id botID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BotState{
				bots:         tt.fields.bots,
				movementList: tt.fields.movementList,
				outputList:   tt.fields.outputList,
			}
			b.move(tt.args.id)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
