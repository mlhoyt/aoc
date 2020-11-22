use rs::intcode::IntCode;

fn main() {
    let input = include_str!("../../input/day02.txt");
    let mut code = input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    // Restore to before "1202 program alarm"
    code[1] = 12;
    code[2] = 2;

    let mut computer = IntCode::new(code);

    if let Err(e) = computer.run() {
        panic!(e);
    }

    println!("{}", computer.mem_at(&0));
}
