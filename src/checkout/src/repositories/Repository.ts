import Container from 'typedi';
import { IRepository } from './IRepository';
import { InMemoryRepository } from './InMemoryRepository';

export function Repository() {
  return function(object: Object, propertyName: string, index?: number) {

    const repository: IRepository = new InMemoryRepository();

    console.log('Creating InMemoryRepository...');

    Container.registerHandler({
      object,
      propertyName,
      index,
      value: () => repository
    });
  };
}