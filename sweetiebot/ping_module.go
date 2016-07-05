package sweetiebot

import (
  "github.com/bwmarrin/discordgo"
)

// This module sucks up all the pings in a message and adds them to the database for the !lastping command
type PingModule struct {
  channels *map[uint64]bool
}

func (w *PingModule) Name() string {
  return "Ping"
}

func (w *PingModule) Register(info *GuildInfo) {
  info.hooks.OnMessageCreate = append(info.hooks.OnMessageCreate, w)
  info.hooks.OnMessageUpdate = append(info.hooks.OnMessageUpdate, w)
}
  
func (w *PingModule)  OnMessageCreate(info *GuildInfo, m *discordgo.Message) {
  w.OnMessageUpdate(info, m)
}

func SBAddPings(m *discordgo.Message) {
  id := SBatoi(m.ID)
  for _, v := range m.Mentions {
    sb.db.AddPing(id, SBatoi(v.ID))
  }
}

func (w *PingModule)  OnMessageUpdate(info *GuildInfo, m *discordgo.Message) {
  SBAddPings(m)
}