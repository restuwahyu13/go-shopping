import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('product_config')
    if (!tablExist) {
      await queryInterface.createTable(
        'product_config',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          name: { type: DataTypes.STRING(200), allowNull: false },
          active: { type: DataTypes.BOOLEAN, allowNull: false },
          promotion_rules: { type: DataTypes.JSONB },
          product_id: { type: DataTypes.UUID, references: { model: 'product', key: 'id' } },
          product_item_id: { type: DataTypes.UUID, references: { model: 'product_item', key: 'id' } },
          min_amount: { type: DataTypes.INTEGER, defaultValue: 0 },
          max_amount: { type: DataTypes.BIGINT, defaultValue: 0 },
          expired_at: { type: DataTypes.DATE },
          created_at: { type: DataTypes.DATE, allowNull: false, defaultValue: Sequelize.literal('CURRENT_TIMESTAMP') },
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
        queryInterface.addIndex('product_config', ['id', 'product_id', 'product_item_id', 'deleted_at'], {
          name: 'product_config_id_product_id_deleted_at_idx',
          unique: true,
          logging: true,
        }),
        queryInterface.addIndex('product_config', ['active', 'min_amount', 'max_amount', 'expired_at'], {
          name: 'product_config_min_amount_max_amount_expired_at_idx',
          unique: true,
          logging: true,
        }),
      ])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('product_config')
    if (tableExist) {
      return queryInterface.dropTable('product_config', { logging: true })
    }
  },
}
