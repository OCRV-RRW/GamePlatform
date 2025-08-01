// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package repository

import (
	database "gameplatform/internal/database"
	uuid "github.com/google/uuid"
	"time"
)

type GameConverterImpl struct{}

func (c *GameConverterImpl) PlatformGameToGetGame(source *database.PlatformGame) GetGame {
	var repositoryGetGame GetGame
	if source != nil {
		var repositoryGetGame2 GetGame
		repositoryGetGame2.ID = c.uuidUUIDToUuidUUID((*source).ID)
		repositoryGetGame2.Title = (*source).Title
		repositoryGetGame2.Description = (*source).Description
		repositoryGetGame2.Src = (*source).Src
		repositoryGetGame2.Icon = (*source).Icon
		repositoryGetGame = repositoryGetGame2
	}
	return repositoryGetGame
}
func (c *GameConverterImpl) PlatformGamesToGetGames(source []database.PlatformGame) []GetGame {
	var repositoryGetGameList []GetGame
	if source != nil {
		repositoryGetGameList = make([]GetGame, len(source))
		for i := 0; i < len(source); i++ {
			repositoryGetGameList[i] = c.databasePlatformGameToRepositoryGetGame(source[i])
		}
	}
	return repositoryGetGameList
}
func (c *GameConverterImpl) PlatformPreviewToGetPreview(source *database.PlatformPreview) GetPreview {
	var repositoryGetPreview GetPreview
	if source != nil {
		var repositoryGetPreview2 GetPreview
		repositoryGetPreview2.ID = c.uuidUUIDToUuidUUID((*source).ID)
		repositoryGetPreview2.Image = (*source).Image
		if (*source).Video != nil {
			xstring := *(*source).Video
			repositoryGetPreview2.Video = &xstring
		}
		repositoryGetPreview = repositoryGetPreview2
	}
	return repositoryGetPreview
}
func (c *GameConverterImpl) PlatformPreviewsToGetPreviews(source []database.PlatformPreview) []GetPreview {
	var repositoryGetPreviewList []GetPreview
	if source != nil {
		repositoryGetPreviewList = make([]GetPreview, len(source))
		for i := 0; i < len(source); i++ {
			repositoryGetPreviewList[i] = c.databasePlatformPreviewToRepositoryGetPreview(source[i])
		}
	}
	return repositoryGetPreviewList
}
func (c *GameConverterImpl) databasePlatformGameToRepositoryGetGame(source database.PlatformGame) GetGame {
	var repositoryGetGame GetGame
	repositoryGetGame.ID = c.uuidUUIDToUuidUUID(source.ID)
	repositoryGetGame.Title = source.Title
	repositoryGetGame.Description = source.Description
	repositoryGetGame.Src = source.Src
	repositoryGetGame.Icon = source.Icon
	return repositoryGetGame
}
func (c *GameConverterImpl) databasePlatformPreviewToRepositoryGetPreview(source database.PlatformPreview) GetPreview {
	var repositoryGetPreview GetPreview
	repositoryGetPreview.ID = c.uuidUUIDToUuidUUID(source.ID)
	repositoryGetPreview.Image = source.Image
	if source.Video != nil {
		xstring := *source.Video
		repositoryGetPreview.Video = &xstring
	}
	return repositoryGetPreview
}
func (c *GameConverterImpl) uuidUUIDToUuidUUID(source uuid.UUID) uuid.UUID {
	var uuidUUID uuid.UUID
	for i := 0; i < len(source); i++ {
		uuidUUID[i] = source[i]
	}
	return uuidUUID
}

type UserConverterImpl struct{}

func (c *UserConverterImpl) CreateUserToCreateUserParams(source *CreateUser) database.CreateUserParams {
	var databaseCreateUserParams database.CreateUserParams
	if source != nil {
		var databaseCreateUserParams2 database.CreateUserParams
		databaseCreateUserParams2.Name = (*source).Name
		databaseCreateUserParams2.Email = (*source).Email
		databaseCreateUserParams2.Password = ByteToString((*source).Password)
		databaseCreateUserParams2.IsAdmin = (*source).IsAdmin
		databaseCreateUserParams2.VerificationCode = (*source).VerificationCode
		databaseCreateUserParams2.Verified = (*source).Verified
		databaseCreateUserParams2.Birthday = c.pTimeTimeToPTimeTime((*source).Birthday)
		if (*source).Gender != nil {
			xstring := *(*source).Gender
			databaseCreateUserParams2.Gender = &xstring
		}
		databaseCreateUserParams = databaseCreateUserParams2
	}
	return databaseCreateUserParams
}
func (c *UserConverterImpl) PlatformUserToGetUser(source *database.PlatformUser) GetUser {
	var repositoryGetUser GetUser
	if source != nil {
		var repositoryGetUser2 GetUser
		repositoryGetUser2.ID = c.uuidUUIDToUuidUUID2((*source).ID)
		repositoryGetUser2.Name = (*source).Name
		repositoryGetUser2.Email = (*source).Email
		repositoryGetUser2.Password = (*source).Password
		repositoryGetUser2.IsAdmin = (*source).IsAdmin
		repositoryGetUser2.VerificationCode = (*source).VerificationCode
		repositoryGetUser2.Verified = (*source).Verified
		repositoryGetUser2.Birthday = c.pTimeTimeToPTimeTime((*source).Birthday)
		if (*source).Gender != nil {
			xstring := *(*source).Gender
			repositoryGetUser2.Gender = &xstring
		}
		repositoryGetUser2.CreatedAt = TimeToTime((*source).CreatedAt)
		repositoryGetUser = repositoryGetUser2
	}
	return repositoryGetUser
}
func (c *UserConverterImpl) PlatformUsersToGetUsers(source []database.PlatformUser) []GetUser {
	var repositoryGetUserList []GetUser
	if source != nil {
		repositoryGetUserList = make([]GetUser, len(source))
		for i := 0; i < len(source); i++ {
			repositoryGetUserList[i] = c.databasePlatformUserToRepositoryGetUser(source[i])
		}
	}
	return repositoryGetUserList
}
func (c *UserConverterImpl) databasePlatformUserToRepositoryGetUser(source database.PlatformUser) GetUser {
	var repositoryGetUser GetUser
	repositoryGetUser.ID = c.uuidUUIDToUuidUUID2(source.ID)
	repositoryGetUser.Name = source.Name
	repositoryGetUser.Email = source.Email
	repositoryGetUser.Password = source.Password
	repositoryGetUser.IsAdmin = source.IsAdmin
	repositoryGetUser.VerificationCode = source.VerificationCode
	repositoryGetUser.Verified = source.Verified
	repositoryGetUser.Birthday = c.pTimeTimeToPTimeTime(source.Birthday)
	if source.Gender != nil {
		xstring := *source.Gender
		repositoryGetUser.Gender = &xstring
	}
	repositoryGetUser.CreatedAt = TimeToTime(source.CreatedAt)
	return repositoryGetUser
}
func (c *UserConverterImpl) pTimeTimeToPTimeTime(source *time.Time) *time.Time {
	var pTimeTime *time.Time
	if source != nil {
		timeTime := TimeToTime((*source))
		pTimeTime = &timeTime
	}
	return pTimeTime
}
func (c *UserConverterImpl) uuidUUIDToUuidUUID2(source uuid.UUID) uuid.UUID {
	var uuidUUID uuid.UUID
	for i := 0; i < len(source); i++ {
		uuidUUID[i] = source[i]
	}
	return uuidUUID
}
