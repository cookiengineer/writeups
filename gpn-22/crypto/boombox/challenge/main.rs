use rand::Rng;
use rand::rngs::ThreadRng;
use num::integer::gcd;
use num_bigint::BigUint;
use std::convert::TryInto;

fn rndint(rng: &mut ThreadRng, a: &BigUint, b: &BigUint) -> BigUint {
    let range = b - a + 1u32;
    let mut buf = vec![0u8; ((range.bits() + 7) / 8).try_into().unwrap()];
    loop {
        rng.fill(&mut buf[..]);
        let num = BigUint::from_bytes_be(&buf);
        if num < range {
            return a + num;
        }
    }
}

fn compose(n: usize, rng: &mut ThreadRng) -> ((Vec<BigUint>, BigUint, BigUint), Vec<BigUint>) {
    let two = BigUint::from(2u32);
    let ps: Vec<BigUint> = (0..n).map(|i| {
        let lower = (&two.pow(i as u32) - 1u32) * &two.pow(n as u32);
        let upper = &two.pow(i as u32) * &two.pow(n as u32);
        rndint(rng, &lower, &upper)
    }).collect();

    let m_lower = &two.pow((2 * n + 1) as u32) + 1u32;
    let m_upper = &two.pow((2 * n + 2) as u32) - 1u32;
    let m = rndint(rng, &m_lower, &m_upper);

    let tune = rndint(rng, &BigUint::from(2u32), &(m.clone() - 2u32));
    let t = &tune / gcd(tune.clone(), m.clone());

    let mictape: Vec<BigUint> = ps.iter().map(|a| (&t * a) % &m).collect();
    if gcd(t.clone(), m.clone()) != BigUint::from(1u32) {
        return compose(n, rng);
    }
    ((ps, t, m), mictape)
}

fn record(msg: &[bool], mic: &[BigUint]) -> BigUint {
    msg.iter().zip(mic).filter(|(&msg_bit, _)| msg_bit).map(|(_, mic_val)| mic_val.clone()).sum()
}

fn main() {
    let n: usize = 42;
    let mut rng = rand::thread_rng();
    let ((ps, t, m), mic) = compose(n, &mut rng);
    // stop thread
    println!("{:?}", mic);

    let msg_str = "GPNCTF{fake_flag}";
    let msg_bytes = msg_str.as_bytes();
    let msg_bin: Vec<bool> = msg_bytes.iter()
                                      .flat_map(|&byte| format!("{:08b}", byte).chars().map(|c| c == '1').collect::<Vec<bool>>())
                                      .collect();
    for chunk in msg_bin.chunks(n) {
        let mut msg_chunk = chunk.to_vec();
        if msg_chunk.len() < n {
            msg_chunk.resize(n, false);
        }
        let c = record(&msg_chunk, &mic);
        println!("{}", c);
    }
}

