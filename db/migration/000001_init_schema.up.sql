CREATE Table Otp(
    Id SERIAL PRIMARY KEY,
    Recipient VARCHAR(50),
    Thema_message VARCHAR(255),
    Body_message VARCHAR(255),
    Date_of_sending DATE
)