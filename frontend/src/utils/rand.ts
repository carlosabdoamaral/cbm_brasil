export function getRandomIntInRange(min: number = 0, max: number): number {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

export function getRandomDate(): Date {
  const start = new Date(2012, 0, 1);
  const end = new Date();

  return new Date(
    start.getTime() + Math.random() * (end.getTime() - start.getTime())
  );
}

export function getRandomBool(): Boolean {
  return getRandomIntInRange(0, 1) % 2 === 0;
}
