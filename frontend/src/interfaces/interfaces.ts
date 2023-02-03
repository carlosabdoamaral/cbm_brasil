export interface NewOccurrenceRequestInterface {
  id_author: number;
  title: string;
  description: string;
  banner_x64: string;
  location: Location; // CEP
}

export interface OccurrenceDetailsInterface {
  IdOccurrence: number;
  Title: string;
  Description: string;
  BannerX64: string;
  IdFirefighter: number;
  FirefighterFullname: string;
  FirefighterEmail: string;
  IdAuthor: number;
  AuthorFullname: string;
  AuthorEmail: string;
  CreatedAt: Date;
  UpdatedAt: Date;
  AcceptedAt: Date;
  CanceledAt: Date;
  FinishedAt: Date;
  Location: Location;
}

export interface Location {
  id: number;
  cep: string;
  country: string;
  state: string;
  city: string;
  neighborhood: string;
  street: string;
  place_number: number;
  complement: string;
  latitude: number;
  longitude: number;
}

export interface SignInInterface {
  Cpf: string;
  Password: string;
}

export interface AddressInterface {
  CEP: string;
  Country: string; // By default it's Brazil - Disabled
  State: string; // Will be filled by an API
  City: string; // Will be filled by an API
  Neighborhood: string; // Will be filled by an API
  Street: string; // Will be filled by an API
  Ddd: string; // Will be filled by an API
  PlaceNumber: number;
  PlaceNotes: string;
}

export interface SignUpInterface {
  Cpf: string;
  FullName: string;
  Age: string;
  MainAddress: AddressInterface;
  Password: string;
}
