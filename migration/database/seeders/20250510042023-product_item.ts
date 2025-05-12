import { QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const products: Record<string, any>[] = await queryInterface.select(null, 'product')

    if (!products?.length) {
      throw new Error('Product not found')
    }

    const productItems: Record<string, any>[] = [
      {
        product_id: products[0]?.id,
        name: 'Google Home',
        sku: '120P90',
        qty: 10,
        sub_brand: 'google home',
        variant: 'black color with super bass E3451',
        category: 'smarthome',
        serial_number: 'GGL0475-0007821',
        unit: 'new',
        buy_amount: 627103,
        sell_amount: 827103,
        ready: true,
      },
      {
        product_id: products[1]?.id,
        name: 'MacBook Pro',
        sku: '43N23P',
        qty: 6,
        sub_brand: 'MacBook',
        variant: 'gray color m1 pro max 512GB',
        category: 'laptop',
        serial_number: 'APL5123-0003874',
        unit: 'new',
        buy_amount: 70345046,
        sell_amount: 89345046,
        ready: true,
      },
      {
        product_id: products[2]?.id,
        name: 'Alexa Speaker Ultra Bass',
        sku: 'A304SD',
        qty: 10,
        sub_brand: 'alexa',
        variant: 'super bass ultra E43113',
        category: 'smart home',
        serial_number: 'AMZ8921-0004182',
        unit: 'new',
        buy_amount: 1000000,
        sell_amount: 1200000,
        ready: true,
      },
      {
        product_id: products[3]?.id,
        name: 'Rasberry Pi B',
        sku: '234234',
        qty: 2,
        sub_brand: 'pi',
        variant: 'mini 512GB E6781',
        category: 'microcontroller',
        serial_number: 'RAB4782-0009867',
        unit: 'new',
        buy_amount: 396363,
        sell_amount: 496363,
        ready: true,
      },
      {
        product_id: products[2]?.id,
        name: 'Alexa Speaker Mini Compo',
        sku: '5A021X',
        qty: 10,
        sub_brand: 'alexa',
        variant: 'mini bass ultra E43113',
        category: 'smart home',
        serial_number: 'AMZ8921-0007891',
        unit: 'new',
        buy_amount: 2000000,
        sell_amount: 3200000,
        ready: true,
      },
    ]
    console.log(productItems)

    return queryInterface.bulkInsert('product_item', productItems, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('product_item', { logging: true })
  },
}
