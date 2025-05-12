import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('courier')
    if (!tablExist) {
      await queryInterface.createTable(
        'courier',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          name: { type: DataTypes.STRING(200), unique: true, allowNull: false },
          code: { type: DataTypes.STRING(25), unique: true, allowNull: false },
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

      return Promise.all([queryInterface.addIndex('courier', ['id', 'code', 'active', 'deleted_at'], { name: 'courier_id_code_active_deleted_at_idx', unique: true, logging: true })])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('courier')
    if (tableExist) {
      return queryInterface.dropTable('courier', { logging: true })
    }
  },
}
