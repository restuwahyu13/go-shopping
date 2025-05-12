import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('bank')
    if (!tablExist) {
      await queryInterface.createTable(
        'bank',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          name: { type: DataTypes.STRING(100), unique: true, allowNull: false },
          code: { type: DataTypes.STRING(25), unique: true, allowNull: false },
          type: { type: DataTypes.STRING(25), allowNull: false },
          active: { type: DataTypes.BOOLEAN, allowNull: false, defaultValue: true },
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

      return Promise.all([queryInterface.addIndex('bank', ['id', 'code', 'active', 'deleted_at'], { name: 'bank_id_code_active_deleted_at_idx', unique: true, logging: true })])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('bank')
    if (tableExist) {
      return queryInterface.dropTable('bank', { logging: true })
    }
  },
}
