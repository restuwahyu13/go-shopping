import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('product_item')
    if (!tablExist) {
      await queryInterface.createTable(
        'product_item',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          product_id: { type: DataTypes.UUID, allowNull: false, references: { model: 'product', key: 'id' } },
          name: { type: DataTypes.STRING(200), allowNull: false },
          sku: { type: DataTypes.STRING(50), unique: true, allowNull: false },
          qty: { type: DataTypes.INTEGER, allowNull: false, defaultValue: 0 },
          sub_brand: { type: DataTypes.STRING(100), allowNull: false },
          category: { type: DataTypes.STRING(100), allowNull: false },
          variant: { type: DataTypes.STRING(100), allowNull: false },
          serial_number: { type: DataTypes.STRING(100), unique: true, allowNull: false },
          unit: { type: DataTypes.STRING(25), allowNull: false },
          buy_amount: { type: DataTypes.BIGINT({ unsigned: true }), allowNull: false },
          sell_amount: { type: DataTypes.BIGINT({ unsigned: true }), allowNull: false },
          ready: { type: DataTypes.BOOLEAN, allowNull: false, defaultValue: false },
          description: { type: DataTypes.TEXT },
          created_at: { type: DataTypes.DATE, allowNull: false, defaultValue: sequelize.literal('CURRENT_TIMESTAMP') },
          created_by: { type: DataTypes.STRING(200) },
          updated_at: { type: DataTypes.DATE },
          updated_by: { type: DataTypes.STRING(200) },
          deleted_at: { type: DataTypes.DATE },
          deleted_by: { type: DataTypes.STRING(200) },
        },
        {
          logging: true,
        },
      )

      return Promise.all([
        queryInterface.addIndex('product_item', ['id', 'product_id', 'deleted_at'], { name: 'product_item_id_product_id_deleted_at_idx', unique: true, logging: true }),
        queryInterface.addIndex('product_item', ['sku', 'qty', 'serial_number', 'unit', 'ready'], { name: 'product_item_sku_qty_serial_number_unit_ready_idx', unique: true, logging: true }),
      ])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('product_item')
    if (tableExist) {
      return queryInterface.dropTable('product_item', { logging: true })
    }
  },
}
