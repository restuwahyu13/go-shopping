import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('product')
    if (!tablExist) {
      await queryInterface.createTable(
        'product',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          brand: { type: DataTypes.STRING(25), unique: true, allowNull: false },
          code: { type: DataTypes.STRING(25), unique: true, allowNull: false },
          active: { type: DataTypes.BOOLEAN, allowNull: false },
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

      return Promise.all([queryInterface.addIndex('product', ['id', 'code', 'active', 'deleted_at'], { name: 'product_id_code_active_deleted_at_idx', unique: true, logging: true })])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('product')
    if (tableExist) {
      return queryInterface.dropTable('product', { logging: true })
    }
  },
}
