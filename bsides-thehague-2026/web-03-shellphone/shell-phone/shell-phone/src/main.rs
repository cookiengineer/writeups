// TODO change everything
use actix_web::{web, App, HttpResponse,  HttpServer, Responder};
use serde::Deserialize;
use std::{env, fs};
use actix_files as fss;
use std::io::{Read, Write};
use std::net::TcpStream;
#[derive(Deserialize)]
struct PhoneData {
    number: String,
}

fn content_type(filename: &str) -> &'static str {
    match filename {
        x if x.ends_with(".html") => "text/html",
        x if x.ends_with(".css") => "text/css",
        x if x.ends_with(".js") => "application/javascript",
        x if x.ends_with(".avif") => "image/avif",
        x if x.ends_with(".jpg") => "image/jpeg",
        x if x.ends_with(".jpeg") => "image/jpeg",
        _ => "text/plain"
    }
}

// Serve static files
async fn serve_file(filename: &str) -> impl Responder {
    let type_header = content_type(filename);
    match fs::read_to_string(filename) {
        Ok(contents) => HttpResponse::Ok().content_type(type_header).body(contents),
        Err(_) => HttpResponse::NotFound().body("File not found"),
    }
}

async fn index() -> impl Responder {
    serve_file("static/index.html").await
}

async fn docs() -> impl Responder {
    serve_file("static/docs.html").await
}

// Handle POST request
async fn dial(form: web::Form<PhoneData>) -> Result<impl Responder, actix_web::Error>{
    let username = env::var("USERNAME").map_err(|_| actix_web::error::ErrorInternalServerError("USERNAME not set"))?;
    let password = env::var("PASSWORD").map_err(|_| actix_web::error::ErrorInternalServerError("PASSWORD not set"))?;
    let asterisk = env::var("ASTERISK").map_err(|_| actix_web::error::ErrorInternalServerError("ASTERISK not set"))?;

    let stream = TcpStream::connect(asterisk.to_string() + ":5038").map_err(|_| actix_web::error::ErrorInternalServerError("Failed to connect to Asterisk server"));

    let connect_action = format!(
        "Action: Login\r\nUsername: {}\r\nSecret: {}\r\n\r\n",
        username, password
    );
    let dial_action = format!("Action: Originate\r\nChannel: {}\r\n\r\n", form.number);

    let mut wstream = match stream {
        Ok(stream) => stream,
        Err(e) => return Err(actix_web::error::ErrorInternalServerError(e.to_string())),
    };
    wstream.write_all(connect_action.as_bytes())?;
    wstream.write_all(dial_action.as_bytes())?;
    wstream.set_read_timeout(Some(std::time::Duration::from_secs(1)))?;
    wstream.flush()?;
    let mut buffer = Vec::new();
    wstream.read_to_end(&mut buffer)?;
    Ok(HttpResponse::Ok().content_type("text/html").body("ok"))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let path = env::current_dir()?;
    println!("Current directory: {:?}", path.display());
    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(index))
            .route("/docs", web::get().to(docs))
            .route("/api/dial", web::post().to(dial))
            .service(fss::Files::new("/static", "./static"))
            .service(fss::Files::new("/webfonts", "./webfonts"))
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}

