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
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
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
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
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

func TestProgram_jumpNZ(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
	}
	type args struct {
		reg   string
		delta string
	}

	one := Register(1)
	zero := Register(0)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"jnz c 2",
			fields{
				cReg: &one,
			},
			args{
				reg:   "c",
				delta: "2"},
			true,
		},
		{
			"jnz c 2",
			fields{
				cReg: &zero,
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
			got := p.jumpNZ(tt.args.reg, tt.args.delta)
			if got != tt.want {
				t.Errorf("Program.jumpNZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_jump(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
	}
	type args struct {
		delta string
	}
	zero_one := Register(0)
	zero_two := Register(0)
	zero_three := Register(0)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"jnz c 2: counter=2",
			fields{counter: 0, cReg: &zero_one},
			args{"2"},
		},
		{
			"jnz 1 2: counter=2",
			fields{counter: 0, cReg: &zero_two},
			args{"1"},
		},
		{
			"jnz c 2: counter=2",
			fields{counter: 0, cReg: &zero_three},
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

func TestProgram_getRegister(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
	}
	type args struct {
		reg string
	}

	aReg := Register(1)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Register
	}{
		{
			"a",
			fields{aReg: &aReg},
			args{"a"},
			&aReg,
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
			if got := p.getRegister(tt.args.reg); got != tt.want {
				t.Errorf("Program.getRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_copy(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
	}
	type args struct {
		from string
		to   string
	}
	one := Register(1)
	two := Register(2)
	another := Register(324)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Register
	}{
		{
			"cpy 1 a",
			fields{aReg: &one},
			args{from: "1", to: "a"},
			Register(1),
		},
		{
			"cpy a b (a=2)",
			fields{aReg: &two, bReg: &another},
			args{from: "a", to: "b"},
			Register(2),
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
			if got := p.copy(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Program.copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_increment(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
	}
	type args struct {
		register string
	}
	one := Register(1)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Register
	}{
		{
			"inc a (1 -> 2)",
			fields{aReg: &one},
			args{"a"},
			Register(2),
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
			if got := p.increment(tt.args.register); got != tt.want {
				t.Errorf("Program.increment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_decrement(t *testing.T) {
	type fields struct {
		counter      int
		instructions []string
		aReg         *Register
		bReg         *Register
		cReg         *Register
		dReg         *Register
	}
	type args struct {
		register string
	}
	two := Register(2)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Register
	}{
		{
			"dec c (2 -> 1)",
			fields{cReg: &two},
			args{"c"},
			Register(1),
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
			if got := p.decrement(tt.args.register); got != tt.want {
				t.Errorf("Program.decrement() = %v, want %v", got, tt.want)
			}
		})
	}
}
