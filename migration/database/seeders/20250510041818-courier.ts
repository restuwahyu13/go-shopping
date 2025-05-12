import { QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const couriers: Record<string, any>[] = [
      {
        name: 'Tiki Jalur Nugraha Ekakurir',
        code: 'JNE',
      },
      {
        name: 'Global Jet Express',
        code: 'J&T',
      },

      {
        name: 'SiCepat Ekspres Indonesia',
        code: 'SICEPAT EXPRESS',
      },
      {
        name: 'Pos Indonesia',
        code: 'POS',
      },
      {
        name: 'Andiarta Muzizat',
        code: 'NINJA XPRESS',
      },
    ]
    console.log(couriers)

    return queryInterface.bulkInsert('courier', couriers, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('courier', { logging: true })
  },
}
