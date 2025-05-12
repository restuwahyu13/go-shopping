import { randomUUID } from 'node:crypto'
import { QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const productRules: Record<string, any>[] = [
      {
        id: randomUUID(),
        name: 'buy',
        code: 'B',
      },
      {
        id: randomUUID(),
        name: 'free',
        code: 'F',
      },
      {
        id: randomUUID(),
        name: 'discount',
        code: 'D',
      },
      {
        id: randomUUID(),
        name: 'PRICE',
        code: 'PR',
      },
      {
        id: randomUUID(),
        name: 'product',
        code: 'P',
      },
      {
        id: randomUUID(),
        name: 'gift',
        code: 'G',
      },
      {
        id: randomUUID(),
        name: 'fixed',
        code: 'FX',
      },
      {
        id: randomUUID(),
        name: 'percentage',
        code: 'PC',
        aritmetic: '%',
      },
      {
        id: randomUUID(),
        name: 'equal',
        code: 'EQ',
        aritmetic: '=',
      },
      {
        id: randomUUID(),
        name: 'not equal',
        code: 'NEQ',
        aritmetic: '!=',
      },
      {
        id: randomUUID(),
        name: 'greather than',
        code: 'GT',
        aritmetic: '>',
      },
      {
        id: randomUUID(),
        name: 'greather than equal',
        code: 'GTE',
        aritmetic: '>=',
      },
      {
        id: randomUUID(),
        name: 'less than',
        code: 'LT',
        aritmetic: '<',
      },
      {
        id: randomUUID(),
        name: 'less than equal',
        code: 'LTE',
        aritmetic: '<=',
      },
      {
        id: randomUUID(),
        name: 'subtracted',
        code: 'SU',
        aritmetic: '-',
      },
      {
        id: randomUUID(),
        name: 'added',
        code: 'AD',
        aritmetic: '+',
      },
      {
        id: randomUUID(),
        name: 'multiplied',
        code: 'MU',
        aritmetic: '*',
      },
      {
        id: randomUUID(),
        name: 'divided',
        code: 'DI',
        aritmetic: '/',
      },
      {
        id: randomUUID(),
        name: 'ALL IN',
        code: 'AL',
      },
      {
        id: randomUUID(),
        name: 'PRODUCT ID',
        code: 'PD',
      },
    ]
    console.log(productRules)

    return queryInterface.bulkInsert('promotion_rules', productRules, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('promotion_rules', { logging: true })
  },
}

/**
 * Example Rules:
 *
 * B(3):EQ:F(1) | { B: 3, EQ: '=', F: 1 } -> Beli 3 produk gratis 1
 * B(2):GT:F(1) | { B: 2, GT: '>', F: 1 } -> Beli lebih dari 2 produk gratis 1
 * B(2):GTE:F(1) | { B: 2, GT: '>', F: 1 } -> Beli 2 produk aatau lebih dari 2 produk gratis 1
 * B(2):GTE:F(1) | { B: 2, GT: '>=', F: 1 } -> Beli 2 produk aatau lebih dari 2 produk gratis 1
 * B(3):EQ:PR(2) | { B: 3, EQ: '=', PR: 2 } -> Beli 3 produk bayar hanya 2 harga
 * B(3):GT:D:PC(10) | { B: 3, GT: '>', D: 10, PC: '%' } ->  Beli lebih dari 3 produk diskon 10%
}
 */

// const x = [
//   // contoh: beli 2 product gratis 1
//   ['B', 'EQ', 'F'],
//   ['B', 'GT', 'F'],
//   ['B', 'GTE', 'F'],
//   ['B', 'LT', 'F'],
//   ['B', 'LTE', 'F'],

//   // contoh: beli 2 product gratis product lainnya
//   ['B', 'EQ', 'F', 'P'],
//   ['B', 'GT', 'F', 'P'],
//   ['B', 'GTE', 'F', 'P'],
//   ['B', 'LT', 'F', 'P'],
//   ['B', 'LTE', 'F', 'P'],

//   // contoh: beli 2 product diskon 10%
//   ['B', 'EQ', 'D', 'PC'],
//   ['B', 'GT', 'D', 'PC'],
//   ['B', 'GTE', 'D', 'PC'],
//   ['B', 'LT', 'D', 'PC'],

//   // contoh: beli 2 product diskon potongan 100RB
//   ['B', 'EQ', 'D', 'SU', 'FX'],
//   ['B', 'GT', 'D', 'SU', 'FX'],
//   ['B', 'GTE', 'D', 'SU', 'FX'],
//   ['B', 'LT', 'D', 'SU', 'FX'],

//   // contoh: beli 3 produk bayar hanya 2 harga
//   ['B', 'EQ', 'PR'],
//   ['B', 'GT', 'PR'],
//   ['B', 'GTE', 'PR'],
//   ['B', 'LT', 'PR'],
// ]
