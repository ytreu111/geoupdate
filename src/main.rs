use std::process::Command;

fn main() {
    let geoip_url = "https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geoip.dat";
    let geosite_url = "https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geosite.dat";

    std::fs::create_dir_all("/tmp/geo-xray").unwrap();

    Command::new("wget")
        .arg("-O")
        .arg("/tmp/geo-xray/geoip.dat")
        .arg(geoip_url)
        .status()
        .expect("failed to execute geoip wget");

    Command::new("wget")
        .arg("-O")
        .arg("/tmp/geo-xray/geosite.dat")
        .arg(geosite_url)
        .status()
        .expect("failed to execute geoip wget");


}
