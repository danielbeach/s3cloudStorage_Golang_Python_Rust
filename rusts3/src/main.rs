use aws_sdk_s3 as s3;
use aws_config::meta::region::RegionProviderChain;
use std::env;
use std::time::Instant;

#[tokio::main]
async fn main() -> Result<(), s3::Error> {
    let now = Instant::now();
    let key = "AWS_ACCESS_KEY_ID";
    env::set_var(key, "");
    let secret = "AWS_SECRET_ACCESS_KEY";
    env::set_var(secret, "");

    let region_provider = RegionProviderChain::default_provider().or_else("us-east-1");
    let config = aws_config::from_env().region(region_provider).load().await;
    let client = s3::Client::new(&config);
    let objects = client.list_objects_v2().bucket("confessions-of-a-data-guy").prefix("202201").send().await?;
    println!("{:?}", objects);
    let elapsed = now.elapsed();
    println!("Elapsed: {:.2?}", elapsed);
Ok(())
}

fn set_evirons() {
    println!("Another function.");
}
