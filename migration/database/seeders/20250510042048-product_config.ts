import { QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const productItems: Record<string, any>[] = await queryInterface.select(null, 'product_item')

    if (!productItems?.length) {
      throw new Error('Product item not found')
    }

    const productConfigs: Record<string, any>[] = [
      {
        name: 'Beli MacBook Pro gratis Rasberry Pi B',
        promotion_rules: JSON.stringify([
          { key: 'B', value: 1 },
          { key: 'EQ', value: '=' },
          { key: 'F', value: productItems[3]?.id },
        ]),
        active: true,
        product_item_id: productItems[1]?.id,
      },
      {
        name: 'Beli 3 Google Home hanya bayar 2 harga',
        promotion_rules: JSON.stringify([
          { key: 'B', value: 3 },
          { key: 'EQ', value: '=' },
          { key: 'PR', value: 2 },
        ]),
        active: true,
        product_item_id: productItems[0]?.id,
      },
      {
        name: 'Beli 3 Alexa Speaker dapatkan discount 10% di semua Alexa Speaker',
        promotion_rules: JSON.stringify([
          { key: 'B', value: 3 },
          { key: 'EQ', value: '=' },
          { key: 'D', value: '10' },
          { key: 'PC', value: '%' },
          { key: 'P', value: 'SAME_BRAND_SIMILAR_CATEGORY' },
        ]),
        active: true,
        product_item_id: productItems[2]?.id,
      },
    ]
    console.log(productConfigs)

    return queryInterface.bulkInsert('product_config', productConfigs, { logging: true })
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    return queryInterface.bulkDelete('product_config', { logging: true })
  },
}
