type Result<T> = std::result::Result<T, IntCodeError>;

#[derive(Debug, Clone)]
pub struct IntCodeError;

impl std::fmt::Display for IntCodeError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "invalid instruction code")
    }
}

pub struct IntCode {
    mem: Vec<i32>,
    pc: usize,
}

impl IntCode {
    pub fn new(code: Vec<i32>) -> IntCode {
        IntCode { mem: code, pc: 0 }
    }

    pub fn run(&mut self) -> Result<bool> {
        loop {
            match self.step() {
                Ok(running) => {
                    if !running {
                        return Ok(false);
                    }
                }
                Err(e) => {
                    return Err(e);
                }
            };
        }
    }

    fn step(&mut self) -> Result<bool> {
        let inst_code = InstrType::new(self.get_mem_at(&self.pc, AddrType::IMMEDIATE));
        match inst_code {
            Some(InstrType::ADD) => {
                let op1 = self.get_mem_at(&(self.pc + 1), AddrType::INDIRECT);
                let op2 = self.get_mem_at(&(self.pc + 2), AddrType::INDIRECT);
                let dst = self.get_mem_at(&(self.pc + 3), AddrType::IMMEDIATE);
                let result = op1 + op2;
                self.set_mem_at(&(dst as usize), result);
                self.pc += 4;
                return Ok(true);
            }
            Some(InstrType::MUL) => {
                let op1 = self.get_mem_at(&(self.pc + 1), AddrType::INDIRECT);
                let op2 = self.get_mem_at(&(self.pc + 2), AddrType::INDIRECT);
                let dst = self.get_mem_at(&(self.pc + 3), AddrType::IMMEDIATE);
                let result = op1 * op2;
                self.set_mem_at(&(dst as usize), result);
                self.pc += 4;
                return Ok(true);
            }
            Some(InstrType::HALT) => {
                self.pc += 1;
                return Ok(false);
            }
            None => {
                return Err(IntCodeError);
            }
        }
    }

    fn get_mem_at(&self, n: &usize, lut: AddrType) -> i32 {
        match lut {
            AddrType::IMMEDIATE => self.mem[*n],
            AddrType::INDIRECT => {
                let nn = self.mem[*n] as usize;
                self.mem[nn]
            }
        }
    }

    fn set_mem_at(&mut self, n: &usize, v: i32) {
        self.mem[*n] = v
    }

    pub fn mem_at(&self, n: &usize) -> i32 {
        self.get_mem_at(n, AddrType::IMMEDIATE)
    }
}

enum InstrType {
    ADD,
    MUL,
    HALT,
}

impl InstrType {
    fn new(v: i32) -> Option<InstrType> {
        match v {
            1 => Some(InstrType::ADD),
            2 => Some(InstrType::MUL),
            99 => Some(InstrType::HALT),
            _ => None,
        }
    }
}

enum AddrType {
    IMMEDIATE,
    INDIRECT,
}

#[test]
fn test_run_add() {
    let code: Vec<i32> = vec![1, 0, 0, 0, 99];
    let expected: Vec<i32> = vec![2, 0, 0, 0, 99];

    let mut computer = IntCode::new(code);
    let still_running = computer.run().unwrap();

    assert_eq!(still_running, false);
    for (i, v) in expected.iter().enumerate() {
        assert_eq!(computer.mem_at(&i), *v);
    }
}

#[test]
fn test_run_mul_simple() {
    let code: Vec<i32> = vec![2, 3, 0, 3, 99];
    let expected: Vec<i32> = vec![2, 3, 0, 6, 99];

    let mut computer = IntCode::new(code);
    let still_running = computer.run().unwrap();

    assert_eq!(still_running, false);
    for (i, v) in expected.iter().enumerate() {
        assert_eq!(computer.mem_at(&i), *v);
    }
}

#[test]
fn test_run_mul_complex() {
    let code: Vec<i32> = vec![2, 4, 4, 5, 99, 0];
    let expected: Vec<i32> = vec![2, 4, 4, 5, 99, 9801];

    let mut computer = IntCode::new(code);
    let still_running = computer.run().unwrap();

    assert_eq!(still_running, false);
    for (i, v) in expected.iter().enumerate() {
        assert_eq!(computer.mem_at(&i), *v);
    }
}

#[test]
fn test_run_add_mul_simple() {
    let code: Vec<i32> = vec![1, 1, 1, 4, 99, 5, 6, 0, 99];
    let expected: Vec<i32> = vec![30, 1, 1, 4, 2, 5, 6, 0, 99];

    let mut computer = IntCode::new(code);
    let still_running = computer.run().unwrap();

    assert_eq!(still_running, false);
    for (i, v) in expected.iter().enumerate() {
        assert_eq!(computer.mem_at(&i), *v);
    }
}

#[test]
fn test_run_add_mul_complex() {
    let code: Vec<i32> = vec![1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50];
    let expected: Vec<i32> = vec![3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50];

    let mut computer = IntCode::new(code);
    let still_running = computer.run().unwrap();

    assert_eq!(still_running, false);
    for (i, v) in expected.iter().enumerate() {
        assert_eq!(computer.mem_at(&i), *v);
    }
}
