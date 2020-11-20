fn main() {
    let input = include_str!("../../input/day01.txt");
    let modules = input
        .lines()
        .map(|line| line.parse::<i32>().unwrap())
        .collect::<Vec<_>>();

    let module_fuel = modules.iter().map(calculate_fuel).fold(0, |sum, x| sum + x);
    println!("module fuel: {}", module_fuel);
}

fn calculate_fuel(mass: &i32) -> i32 {
    let fuel = (((*mass as f32) / 3.0).floor() as i32) - 2;
    clamp_min(&fuel, &0)
}

#[test]
fn test_calculate_fuel() {
    assert_eq!(calculate_fuel(&12), 2);
    assert_eq!(calculate_fuel(&14), 2);
    assert_eq!(calculate_fuel(&1969), 654);
    assert_eq!(calculate_fuel(&100756), 33583);
}

fn clamp_min<T>(v: &T, limit: &T) -> T
where
    T: Ord + Copy,
{
    if v < limit {
        return *limit;
    }

    return *v;
}
