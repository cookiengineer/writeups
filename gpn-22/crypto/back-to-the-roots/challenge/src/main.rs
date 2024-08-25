use std::fs::File;
use std::io::Read;
use rand::Rng;

struct Grid<T: Copy> {
    /// row major, may be shorter than width * height
    pub data: Vec<T>,
    pub width: usize,
    pub height: usize,
}

impl<T: Copy> Grid<T> {
    pub fn column(&self, column: usize) -> Vec<T> {
        let mut result = Vec::new();
        for row in 0..self.height {
            let index = self.index(row, column);
            if index >= self.data.len() {
                continue;
            }
            result.push(self.data[index]);
        }
        result
    }

    pub fn at(&self, row: usize, column: usize) -> T {
        self.data[self.index(row, column)]
    }

    pub fn index(&self, row: usize, column: usize) -> usize {
        row * self.width + column
    }
}

fn encrypt_message(key: &str, msg: &str) -> String {
    //I was told that it would be really hard to break if you just used bits instead of letters
    let message:String = msg.to_string().bytes()
            .flat_map(|b| format!("{:08b}", b).chars().collect::<Vec<char>>())
            .collect();
    let mut grid = Grid {
        data: Vec::new(),
        width: key.len(),
        height: message.len().div_ceil(key.len()),
    };

    message.chars().for_each(|c| grid.data.push(c));

    let mut sorted_key: Vec<(usize, char)> = (0..key.len()).zip(key.chars()).collect();
    sorted_key.sort_by_key(|pair| pair.1);

    let mut result = String::new();

    for (index, _) in sorted_key {
        result.extend(grid.column(index))
    }

    result
}
fn genKey()->String{
    let key_len = rand::thread_rng().gen_range(21..35);
    const CHARSET: &[u8] = b"abcdefghijklmnopqrstuvwxyz";
    let mut rng = rand::thread_rng();
    let key: String = (0..key_len)
        .map(|_| {
            let idx = rng.gen_range(0..CHARSET.len());
            CHARSET[idx] as char
        })
        .collect();
    key
}
fn main() -> std::io::Result<()> {
    //LPT: You can't leak your keys if you don't know your keys...
    let mut file = File::open("flag.txt")?;
    let mut message= String::new();
    file.read_to_string(&mut message)?;

    let key = genKey();
    message = message.replace("GPNCTF{","");
    message = message.replace("}","");
    message = message.to_ascii_lowercase();
    assert!(message.chars().all(|c| c.is_ascii_alphabetic() && c.is_ascii_lowercase()));
    let out = encrypt_message(&key,&message);
    println!("{:?}",out);
    Ok(())
}
