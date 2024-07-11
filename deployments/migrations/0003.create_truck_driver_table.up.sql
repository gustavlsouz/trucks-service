create table
  truck_driver (
    id uuid primary key,
    truckId uuid not null unique,
    driverId uuid not null,
    createdAt timestamp without time zone not null default now(),
    constraint fk_truck_driver_driverId foreign key (driverId) references public.driver (id),
    constraint fk_truck_driver_truckId foreign key (truckId) references public.truck (id)
  );