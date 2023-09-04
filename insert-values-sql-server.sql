USE POC
GO

CREATE TABLE Employees
(
	ID BIGINT IDENTITY PRIMARY KEY,
	Role VARCHAR(255) NULL, 
	Company VARCHAR(255) NULL,
	Location VARCHAR(255) NULL, 
	Remote BIT NOT NULL,
	Link VARCHAR(255) NULL,
	Salary MONEY NULL,
	CreatedAt DATETIMEOFFSET DEFAULT SYSDATETIMEOFFSET(),
	UpdatedAt DATETIMEOFFSET NULL,
	DeletedAt DATETIMEOFFSET NULL
);

INSERT INTO Employees (Role, Company, Location, Remote, Link, Salary) VALUES ('Senior Software Enginner', 'Neon', 'Goiânia', 1, 'https://linkedin.com', 14500.00);
INSERT INTO Employees (Role, Company, Location, Remote, Link, Salary) VALUES ('Software Enginner III', 'Neon', 'Goiânia', 1, 'https://linkedin.com', 9500.00);
INSERT INTO Employees (Role, Company, Location, Remote, Link, Salary) VALUES ('Software Enginner II', 'Neon', 'Goiânia', 1, 'https://linkedin.com', 5000.00);
INSERT INTO Employees (Role, Company, Location, Remote, Link, Salary) VALUES ('Staff Software Enginner', 'Neon', 'Goiânia', 1, 'https://linkedin.com', 20000.00);
INSERT INTO Employees (Role, Company, Location, Remote, Link, Salary) VALUES ('Distinguished Software Enginner', 'Neon', 'Goiânia', 1, 'https://linkedin.com', 48000.00);
INSERT INTO Employees (Role, Company, Location, Remote, Link, Salary) VALUES ('Fellow Software Enginner', 'Neon', 'Goiânia', 1, 'https://linkedin.com', 70000.00);