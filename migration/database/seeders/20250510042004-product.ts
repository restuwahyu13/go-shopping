import { QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const products: Record<string, any>[] = [
      {
        brand: 'google',
        code: 'GGL0475',
        active: true,
      },
      {
        brand: 'apple',
        code: 'APL5123',
        active: true,
      },
      {
        brand: 'amazon',
        code: 'AMZ8921',
        active: true,
      },
      {
        brand: 'raspberry',
        code: 'RAB4782',
        active: true,
      },
    ]
    console.log(products)

    return queryInterface.bulkInsert('product', products, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('product', { logging: true })
  },
}
