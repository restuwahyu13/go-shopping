import { QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const banks: Record<string, any>[] = [
      {
        name: 'Bank Central Asia',
        code: 'BCA',
        type: 'bank',
      },
      {
        name: 'Bank Mandiri',
        code: 'MANDIRI',
        type: 'bank',
      },
      {
        name: 'Bank Nasional Indonesia',
        code: 'BNI',
        type: 'bank',
      },
      {
        name: 'Bank Rakyat Indonesia',
        code: 'BRI',
        type: 'bank',
      },
      {
        name: 'Bank Artos Indonesia',
        code: 'JAGO',
        type: 'bank',
      },
    ]
    console.log(banks)

    return queryInterface.bulkInsert('bank', banks, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('bank', { logging: true })
  },
}
