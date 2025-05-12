import { QueryInterface, Sequelize, DataTypes } from 'sequelize'

module.exports = {
  up: async (queryInterface: QueryInterface, sequelize: Sequelize) => {
    const tablExist: boolean = await queryInterface.tableExists('user')
    if (!tablExist) {
      await queryInterface.createTable(
        'users',
        {
          id: { type: DataTypes.UUID, primaryKey: true, allowNull: false, unique: true, defaultValue: sequelize.literal('uuid_generate_v4()') },
          name: { type: DataTypes.STRING(200), allowNull: false },
          email: { type: DataTypes.STRING(50), allowNull: false },
          status: { type: DataTypes.STRING(25), allowNull: false, defaultValue: 'active' },
          password: { type: DataTypes.TEXT, allowNull: false },
          verified_at: { type: DataTypes.DATE },
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
        queryInterface.addIndex('users', ['id'], { name: 'users_id_idx', unique: true, logging: true }),
        queryInterface.addIndex('users', ['email', 'status', 'verified_at', 'deleted_at'], { name: 'users_email_status_verified_at_deleted_at_idx', unique: true, logging: true }),
      ])
    }
  },
  down: async (queryInterface: QueryInterface, _sequelize: Sequelize) => {
    const tableExist: boolean = await queryInterface.tableExists('users')
    if (tableExist) {
      return queryInterface.dropTable('users')
    }
  },
}
