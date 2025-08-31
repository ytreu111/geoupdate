use std::process::Command;

fn main() {
    let geoip_url =
        "https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geoip.dat";
    let geosite_url =
        "https://raw.githubusercontent.com/runetfreedom/russia-v2ray-rules-dat/release/geosite.dat";
    let tmp_dir_src = "/tmp/geo-xray";
    let tmp_geoip_src = format!("{}/geoip.dat", tmp_dir_src);
    let tmp_geosite_src = format!("{}/geosite.dat", tmp_dir_src);
    let dest_dir_src = "/usr/share/xray";
    let geoip_dest = format!("{}/geoip.dat", dest_dir_src);
    let geosite_dest = format!("{}/geosite.dat", dest_dir_src);

    std::fs::create_dir_all("/tmp/geo-xray").unwrap();

    Command::new("wget")
        .arg("-O")
        .arg(tmp_geoip_src.clone())
        .arg(geoip_url)
        .status()
        .expect("failed to execute geoip wget");

    Command::new("wget")
        .arg("-O")
        .arg(tmp_geosite_src.clone())
        .arg(geosite_url)
        .status()
        .expect("failed to execute geoip wget");

    std::fs::copy(tmp_geoip_src, geoip_dest).expect("failed to copy geoip");
    std::fs::copy(tmp_geosite_src, geosite_dest).expect("failed to copy geosite");
    std::fs::remove_dir_all(tmp_dir_src).expect("failed to remove dir");
}
