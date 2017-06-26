INSERT INTO User (`Email`, `FirstName`, `LastName`, `Password`) VALUES
  ('n.i.vdveer@gmail.com', 'Niels', 'van der Veer', 'supersecret'),
  ('thom@doge.beer', 'Thom', 'Overhand', 'supersecretdoge');

INSERT INTO Project (`Name`) VALUES
  ('Super awesome App'), ('Declaration App');

INSERT INTO Receipt (`ImagePath`, `Data`) VALUES
  ('path/to/image', ''), ('path/to/image', '');

INSERT INTO Declaration (`Title`, `TotalPrice`, `VATPrice`, `Date`, `Description`, `ProjectID`, `StoreName`, `ReceiptID`, `UserID`) VALUES
  ('Lunch', 19.99, 4.1979, '2017-07-09 12:32:00', 'Lunch voor onderweg', 1, 'Albert Heijn', 1, 1);

INSERT INTO DeclarationStatus (`Status`, `DateModifed`, `DeclarationID`, `ModifiedByUserId`) VALUES
  ('pending', '2017-07-09 12:32:00', 1, 1);

INSERT INTO UserProject VALUES
  (1, 1);