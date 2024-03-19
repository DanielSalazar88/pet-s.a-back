SELECT
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'nombre_mascota',
            m3.nombre,
            'raza_mascota',
            m3.raza,
            'peso_mascota',
            m3.peso,
            'edad_mascota',
            m3.edad,
            'recetas',
            recetas_json
        )
    ) AS resultado
FROM
    Mascota m3
    INNER JOIN (
        SELECT
            r.mascota_id,
            JSON_ARRAYAGG(
                JSON_OBJECT(
                    'medicamento',
                    m4.nombre,
                    'descripcion_receta',
                    m4.descripcion,
                    'dosis_receta',
                    m4.dosis
                )
            ) AS recetas_json
        FROM
            Receta r
            INNER JOIN Medicamento m4 ON r.medicamento_id = m4.id
        GROUP BY
            r.mascota_id
    ) AS recetas_agrupadas ON recetas_agrupadas.mascota_id = m3.id
WHERE
    m3.id = id_mascota_replace;