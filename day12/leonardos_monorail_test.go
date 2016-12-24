package main

import "testing"

func Test_isRegister(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"a (true)",
			args{"a"},
			true,
		},
		{
			"b (true)",
			args{"b"},
			true,
		},
		{
			"c (true)",
			args{"c"},
			true,
		},
		{
			"c (true)",
			args{"d"},
			true,
		},
		{
			"12 (false)",
			args{"12"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRegister(tt.args.val); got != tt.want {
				t.Errorf("isRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_execute(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	type args struct {
		instr []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			if got := p.execute(tt.args.instr); got != tt.want {
				t.Errorf("Program.execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_Run(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			p.Run()
		})
	}
}

func TestProgram_copy(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	type args struct {
		from string
		to   string
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
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			p.copy(tt.args.from, tt.args.to)
		})
	}
}

func TestProgram_increment(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	type args struct {
		register string
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
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			p.increment(tt.args.register)
		})
	}
}

func TestProgram_decrement(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	type args struct {
		register string
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
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			p.decrement(tt.args.register)
		})
	}
}

func TestProgram_jumpNZ(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	type args struct {
		reg   string
		delta string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"jnz c 2",
			fields{
				cReg: Register(1),
			},
			args{
				reg:   "c",
				delta: "2"},
			true,
		},
		{
			"jnz c 2",
			fields{
				cReg: Register(0),
			},
			args{
				reg:   "c",
				delta: "2"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			if got := p.jumpNZ(tt.args.reg, tt.args.delta); got != tt.want {
				t.Errorf("Program.jumpNZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_jump(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         Register
		bReg         Register
		cReg         Register
		dReg         Register
	}
	type args struct {
		delta string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"jnz c 2: counter=2",
			fields{counter: 0, cReg: Register(0)},
			args{"2"},
		},
		{
			"jnz 1 2: counter=2",
			fields{counter: 0, cReg: Register(0)},
			args{"1"},
		},
		{
			"jnz c 2: counter=2",
			fields{counter: 0, cReg: Register(0)},
			args{"2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Program{
				counter:      tt.fields.counter,
				instructions: tt.fields.instructions,
				aReg:         tt.fields.aReg,
				bReg:         tt.fields.bReg,
				cReg:         tt.fields.cReg,
				dReg:         tt.fields.dReg,
			}
			p.jump(tt.args.delta)
		})
	}
}
