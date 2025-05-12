import { DataTypes, QueryInterface, Sequelize } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('promotion_rules')
    if (!tablExist) {
      await queryInterface.createTable(
        'promotion_rules',
        {
          id: { type: DataTypes.UUID, primaryKey: true, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          name: { type: DataTypes.STRING(100), unique: true, allowNull: false },
          code: { type: DataTypes.STRING(15), unique: true, allowNull: false },
          aritmetic: { type: DataTypes.STRING(10) },
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

      return Promise.all([queryInterface.addIndex('promotion_rules', ['id', 'code', 'deleted_at'], { name: 'promotion_rules_id_code_deleted_at_idx', unique: true, logging: true })])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('promotion_rules')
    if (tableExist) {
      return queryInterface.dropTable('promotion_rules', { logging: true })
    }
  },
}
