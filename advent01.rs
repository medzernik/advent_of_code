use std::fs;

fn main() {
    let mut final_vysledok_dve = 0;
    let mut final_vysledok_tri = 0;
    let mut sum:i32;

    let mut v = Vec::new();
    let contents =
        fs::read_to_string("input.txt").expect("Something went wrong: reading file error");

    for s in contents.lines() {
        if !s.is_empty() {
            v.push(s.parse::<i32>().unwrap());
            println!("{}", s);
        }
    }
    println!("--------------");

    for (_i, x) in v.iter().enumerate() {
        //eprintln!("In position {} we have value {}", i, x);
        for (_j, y) in v.iter().enumerate() {
            sum = x + y;
            //eprintln!("Summing {} + {} = {} ", x ,z, sum);
            if sum == 2020 {
                println!("MATCH 2020 DVE CISLA FOUND!!! {} + {} = {}", x, y, sum);
                final_vysledok_dve = x * y;
            }
            for (_k, z) in v.iter().enumerate() {
                sum = x + y + z;
                //eprintln!("Summing {} + {} + {} = {} ", x ,y, z, sum);
                if sum == 2020 {
                    println!(
                        "MATCH 2020 TROJICA FOUND!!! {} + {} + {} = {}",
                        x, y, z, sum
                    );
                    final_vysledok_tri = x * y * z;
                }
            }
        }
    }

    println!("Final vysledok dve cisla: {}", final_vysledok_dve);
    println!("Final vysledok tri cisla: {}", final_vysledok_tri);
}
