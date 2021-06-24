use std::fs as f;

// Location where the information regarding the battery status and
// capacity is stored in Linux systems

const CAPACITY_PATH: &str = "/sys/class/power_supply/BAT0/capacity";
const STATUS_PATH: &str = "/sys/class/power_supply/BAT0/status";

fn main() {
    let contents =
        f::read_to_string(STATUS_PATH).expect("Something went wrong reading the battery status");

    println!("Current battery status: {}", contents);

    let contents = f::read_to_string(CAPACITY_PATH)
        .expect("Something went wrong reading the battery percentage");

    println!("Current battery percentage: {}", contents);
}
