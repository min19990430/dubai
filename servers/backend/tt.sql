SELECT
    time_series.times as times,
    pH_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as pH,
    pH_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as pH_status,
    EC_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as EC,
    EC_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as EC_status,
    COD_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as COD,
    COD_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as COD_status,
    Turbidity_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as Turbidity,
    Turbidity_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as Turbidity_status,
    Temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as Temp,
    Temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as Temp_status,
    battery_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as battery_temp,
    battery_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as battery_temp_status,
    battery_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as battery,
    battery_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as battery_status,
    signal_strength_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as signal_strength,
    signal_strength_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as signal_strength_status,
    FCM_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.value as FCM_temp,
    FCM_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.status as FCM_temp_status
FROM
    (
        SELECT
            datetime as times
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            physical_quantity_uuid = '47d28dc4-b888-4880-ab23-3d1c46b24f06'
            AND (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
        ORDER BY
            datetime
    ) as time_series
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = '47d28dc4-b888-4880-ab23-3d1c46b24f06'
        ORDER BY
            datetime
    ) as pH_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = pH_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = '8df500a9-41b8-4aa0-b072-2fa2313f9677'
        ORDER BY
            datetime
    ) as EC_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = EC_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = 'f17c29aa-9ee0-476b-9b89-4d624fd1af82'
        ORDER BY
            datetime
    ) as COD_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = COD_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = '9019c4e1-e40d-4e77-b69e-87f70c4de96d'
        ORDER BY
            datetime
    ) as Turbidity_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = Turbidity_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = 'c9d30b7c-bcdd-466b-9c75-29dfcf3dc59f'
        ORDER BY
            datetime
    ) as Temp_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = Temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = '184c9a57-65bc-4271-97d6-b983431e511e'
        ORDER BY
            datetime
    ) as battery_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = battery_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = 'a03598ab-8b1a-48ec-8ecf-d1e7838127df'
        ORDER BY
            datetime
    ) as battery_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = battery_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = '675a68b9-e5f2-4b25-9ddb-d73cfb4c2d6a'
        ORDER BY
            datetime
    ) as signal_strength_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = signal_strength_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
    LEFT JOIN (
        SELECT
            datetime as pqtime,
            value,
            status
        FROM
            `e2232881-e996-46be-a0a2-95a2983fdb42`
        WHERE
            (
                datetime BETWEEN '2024-05-08 09:00:00'
                AND '2024-05-08 11:00:00'
            )
            AND physical_quantity_uuid = '5c5d94c5-d1e2-4519-81ed-41ea88def09a'
        ORDER BY
            datetime
    ) as FCM_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182 ON time_series.times = FCM_temp_f52132d2_8a8f_4acc_afbd_6ec556b12182.pqtime
ORDER BY
    times