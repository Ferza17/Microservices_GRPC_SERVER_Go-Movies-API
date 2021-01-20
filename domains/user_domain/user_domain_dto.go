package user_domain

type Wishlist struct {
	IdWishlist int64  `db:"Id_wishlist"`
	IdMovie    string `db:"Id_movie"`
	IdUser     int64  `db:"Id_user"`
}

type Watched struct {
	IdWatched int64   `db:"Id_watched"`
	IdMovie   string  `db:"Id_movies"`
	IdUser    int64   `db:"Id_user"`
	Rate      float32 `db:"Rate"`
}

// TODO : @add Voucher

type User struct {
	Id       int64       `db:"id"`
	Name     string      `db:"Name"`
	Email    string      `db:"Email"`
	Password string      `db:"Password"`
	Phone    string      `db:"Phone"`
	Payment  float32     `db:"Payment"`
	Loyalty  int32       `db:"Loyalty"`
	Wishlist []*Wishlist `db:"Wishlist"`
	Watched  []*Watched  `db:"Watched"`
}
