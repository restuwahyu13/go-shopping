import { QueryInterface, Sequelize } from 'sequelize'
import { faker } from '@faker-js/faker'
import bcrypt from 'bcrypt'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const hashPassword: string = bcrypt.hashSync('@Qwerty12', 12)
    const users: Record<string, any>[] = [
      {
        name: 'Restu Wahyu Saputra',
        email: 'restuwahyu705@gmail.com',
        password: hashPassword,
        verified_at: new Date(),
      },
    ]

    for (let i = 1; i <= 10; i++) {
      users.push({
        name: faker.person.fullName(),
        email: faker.internet.email(),
        password: hashPassword,
        verified_at: new Date(),
      })
    }
    console.log(users)

    return queryInterface.bulkInsert('users', users, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('users', { logging: true })
  },
}
