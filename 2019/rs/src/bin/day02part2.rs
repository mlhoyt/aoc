use rs::intcode::IntCode;

fn main() {
    let input = include_str!("../../input/day02.txt");
    let code = input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    for noun in 0..100 {
        for verb in 0..100 {
            let mut test_code = code.clone();

            // Restore to before "1202 program alarm"
            test_code[1] = noun;
            test_code[2] = verb;

            let mut computer = IntCode::new(test_code);

            if let Err(_) = computer.run() {
                continue;
            }

            if computer.mem_at(&0) == 19690720 {
                println!("{}", 100 * noun + verb);
                return;
            }
        }
    }

    println!("no noun/verb match found");
}
