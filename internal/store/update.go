// Copyright (C) 2022  Aaron González

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package store

func (m *Manager) ToggleStatus(id string) (err error) {
	update, err := m.db.Prepare("UPDATE tasks SET done = ((done | 1) - (done & 1)) WHERE id = ?")
	if err != nil {
		return
	}
	defer update.Close()

	_, err = update.Exec(id)
	if err != nil {
		return
	}

	return
}
