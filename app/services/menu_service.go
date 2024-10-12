package services

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type MenuService struct {
}

func NewMenuService() *MenuService {
	return &MenuService{}
}

type SubMenuProps struct {
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Route string `json:"route"`
}

type SubMenuItems struct {
	Title string       `json:"title"`
	Icon  string       `json:"icon"`
	Props SubMenuProps `json:"props"`
}

type MenuProps struct {
	Type     string         `json:"type"`
	Icon     string         `json:"icon"`
	Route    string         `json:"route"`
	Submenus []SubMenuItems `json:"submenus"`
}

type MenuItem struct {
	Title string    `json:"title"`
	Props MenuProps `json:"props"`
	Route string    `json:"route"`
}

func (s *MenuService) Fetch(ctx http.Context) ([]MenuItem, error) {
	var user models.User

	err := facades.Auth(ctx).User(&user)

	if err != nil {
		return []MenuItem{}, err
	}

	var authent string = user.Authent

	if authent == "superadmin" {
		menus := []MenuItem{
			{
				Title: "Dashboard",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-view-dashboard",
					Route: "/auth/logged/dashboard",
				},
				Route: "/auth/logged/dashboard",
			}, {
				Title: "Aplikasi",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-monitor-dashboard",
				},
			},
			{
				Title: "Data Master",
				Props: MenuProps{
					Type: "group",
					Icon: "mdi-database-settings",
					Submenus: []SubMenuItems{
						{
							Title: "Konfigurasi Aplikasi",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/master-app-info",
							},
						},
					},
				},
			},
			{
				Title: "WEB MPP",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-web",
				},
			},
			{
				Title: "Utilitas",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-tools",
				},
			}, {
				Title: "Manajemen Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-group",
					Route: "/auth/logged/utility-user-manajemen",
				},
			}, {
				Title: "Profil Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-details",
					Route: "/auth/logged/utility-user-profile",
				},
			},
			{
				Title: "Ubah Kata Sandi",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-key",
					Route: "/auth/logged/utility-user-change-pwd",
				},
			},
			{
				Title: "File Management",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-file",
					Route: "/auth/logged/utility-file-management",
				},
			},
		}

		return menus, nil
	}

	if authent == "administrator" {
		menus := []MenuItem{
			{
				Title: "Dashboard",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-view-dashboard",
					Route: "/auth/logged/dashboard",
				},
				Route: "/auth/logged/dashboard",
			}, {
				Title: "Aplikasi",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-monitor-dashboard",
				},
			},
			{
				Title: "Data Master",
				Props: MenuProps{
					Type: "group",
					Icon: "mdi-database-settings",
					Submenus: []SubMenuItems{
						{
							Title: "Manajemen Puskesmas",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/master-puskesmas",
							},
						},
						{
							Title: "Manajemen Posyandu",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/master-posyandu",
							},
						},
					},
				},
			},
			{
				Title: "e-Dasawisma",
				Props: MenuProps{
					Type: "group",
					Icon: "mdi-account-child-circle",
					Submenus: []SubMenuItems{
						{
							Title: "Import Data Pasien Anak",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/edasawisma-import-data-pasien",
							},
						},
						{
							Title: "Data Pasien Anak",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/edasawisma-data-pasien",
							},
						},
						{
							Title: "Data Pemantauan",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/edasawisma-data-pemantauan",
							},
						},

						{
							Title: "Data Penerima Bantuan",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/edasawisma-data-pemantauan-bantuan",
							},
						},
					},
				},
			}, {
				Title: "Utilitas",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-tools",
				},
			}, {
				Title: "Manajemen Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-group",
					Route: "/auth/logged/utility-user-manajemen",
				},
			}, {
				Title: "Profil Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-details",
					Route: "/auth/logged/utility-user-profile",
				},
			},
			{
				Title: "Ubah Kata Sandi",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-key",
					Route: "/auth/logged/utility-user-change-pwd",
				},
			},
		}

		return menus, nil
	}

	//menud umped
	if authent == "enumerator" {
		menus := []MenuItem{
			{
				Title: "Dashboard",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-view-dashboard",
					Route: "/auth/logged/dashboard",
				},
				Route: "/auth/logged/dashboard",
			}, {
				Title: "Aplikasi",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-monitor-dashboard",
				},
			},

			{
				Title: "e-Dasawisma",
				Props: MenuProps{
					Type: "group",
					Icon: "mdi-account-child-circle",
					Submenus: []SubMenuItems{
						{
							Title: "Data Pemantauan",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/edasawisma-data-pemantauan",
							},
						},

						{
							Title: "Data Penerima Bantuan",
							Icon:  "",
							Props: SubMenuProps{
								Type:  "item",
								Route: "/auth/logged/edasawisma-data-pemantauan-bantuan",
							},
						},
					},
				},
			}, {
				Title: "Utilitas",
				Props: MenuProps{
					Type: "subheader",
					Icon: "mdi-tools",
				},
			}, {
				Title: "Profil Pengguna",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-account-details",
					Route: "/auth/logged/utility-user-profile",
				},
			},
			{
				Title: "Ubah Kata Sandi",
				Props: MenuProps{
					Type:  "item",
					Icon:  "mdi-key",
					Route: "/auth/logged/utility-user-change-pwd",
				},
			},
		}

		return menus, nil
	}

	return []MenuItem{}, nil

}
