SELECT JSON_ARRAYAGG(
    JSON_OBJECT(
        'id', id,
        'nombre', nombre,
        'raza', raza,
        'edad', edad,
        'peso', peso,
        'cliente_cedula', cliente_cedula
    )
) AS resultado
FROM Mascota;
