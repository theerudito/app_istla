-- CREATE DATABASE istla;
CREATE TABLE
    IF NOT EXISTS perfiles (
        id_perfil INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        descripcion VARCHAR(150)
    );

CREATE TABLE
    IF NOT EXISTS usuarios (
        id_usuario INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        identificacion VARCHAR(13) UNIQUE,
        nombres VARCHAR(50),
        apellidos VARCHAR(50),
        email VARCHAR(150),
        password VARCHAR(150),
        id_perfil INT NOT NULL,
        FOREIGN KEY (id_perfil) REFERENCES perfiles (id_perfil)
    );

CREATE TABLE
    IF NOT EXISTS storage (
        id_storage INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        nombre VARCHAR,
        url VARCHAR,
        extencion VARCHAR
    );

CREATE TABLE
    IF NOT EXISTS post_usuario (
        id_post_usuario INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        descripcion VARCHAR,
        id_usuario INT NOT NULL,
        id_storage INT DEFAULT NULL,
        usuario_creacion VARCHAR(100),
        usuario_modificacion VARCHAR(100),
        fecha_creacion TIMESTAMP DEFAULT now (),
        fecha_modificacion TIMESTAMP DEFAULT now (),
        FOREIGN KEY (id_usuario) REFERENCES usuarios (id_usuario),
        FOREIGN KEY (id_storage) REFERENCES storage (id_storage)
    );

CREATE TABLE
    IF NOT EXISTS logs_accion (
        id_log INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        accion VARCHAR(100) NOT NULL,
        tabla_nombre VARCHAR(100),
        descripcion TEXT,
        id_registro INT,
        fecha TIMESTAMP DEFAULT now ()
    );

CREATE TABLE
    IF NOT EXISTS logs_error (
        id_log_error INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        mensaje TEXT,
        tabla_nombre VARCHAR(100),
        fecha TIMESTAMP DEFAULT now ()
    );